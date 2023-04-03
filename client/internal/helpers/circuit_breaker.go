package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

type CircuitBreaker struct {
	// Максимальное количество ошибок, которые могут быть допущены,
	// прежде чем Circuit breaker переводится в режим "отключен".
	MaxErrors int
	// Время, в течение которого должны быть допущены MaxErrors ошибок,
	// прежде чем Circuit breaker переводится в режим "отключен".
	Timeout time.Duration
	// Время, которое должно пройти, прежде чем Circuit breaker переводится
	// из режима "отключен" в режим "полуоткрыт".
	ResetTime time.Duration
	// Статус Circuit breaker
	Status string
	// Количество ошибок, допущенных за период Timeout
	ErrorCount int
	// Время последней ошибки
	LastFailure time.Time
}

// Функция для проверки, должен ли Circuit breaker перевестись в режим "отключен".
func (cb *CircuitBreaker) IsOpen() bool {
	if cb.ErrorCount > cb.MaxErrors && time.Now().Sub(cb.LastFailure) < cb.Timeout {
		return true
	}
	return false
}

// Функция для проверки, можно ли снова использовать Circuit breaker.
func (cb *CircuitBreaker) IsClosed() bool {
	if time.Now().Sub(cb.LastFailure) > cb.ResetTime {
		return true
	}
	return false
}

func (cb *CircuitBreaker) ExecuteRequest(url string, c *gin.Context) {
	if cb.IsOpen() {
		c.JSON(http.StatusServiceUnavailable, []map[string]interface{}{})
		return
	}

	resp, err := http.Get(url)

	if err != nil {
		cb.ErrorCount++
		cb.LastFailure = time.Now()
		return
	}

	cb.ErrorCount = 0
	cb.LastFailure = time.Time{}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var data []map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return
	}
	c.JSON(http.StatusOK, data)
	return
}
