package handlers

import (
	"client/internal/helpers"
	"client/internal/infrastructure/env"
	"client/internal/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	services       *services.Service
	client         *http.Client
	circuitBreaker *helpers.CircuitBreaker
}

func NewHandler(services *services.Service, circuitBreaker *helpers.CircuitBreaker) *Handler {
	return &Handler{services: services, circuitBreaker: circuitBreaker}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("/add-film", h.addFilm)
		api.POST("/add-review", h.addReview)
		api.GET("/report", h.report)
		api.GET("/report2", h.report2)
		api.GET("/report3", h.report3)
	}
	return router
}

//func CircuitBreakerMiddleware(cb *helpers.CircuitBreaker, next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		if err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			w.Write([]byte("Circuit breaker tripped"))
//			return
//		}
//
//		next.ServeHTTP(w, r)
//	})
//}

func (h *Handler) addFilm(c *gin.Context) {

	var input services.AddRequest
	if err := c.BindJSON(&input); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	input.RoutingKey = "film"

	err := h.services.FilmService.Create(input)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"status": "ok",
	})
}

func (h *Handler) addReview(c *gin.Context) {

	var input services.AddRequest
	if err := c.BindJSON(&input); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	input.RoutingKey = "review"

	err := h.services.FilmService.Create(input)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"status": "ok",
	})
}

func (h *Handler) report(c *gin.Context) {
	h.circuitBreaker.ExecuteRequest(fmt.Sprintf(env.DataServiceHost, "report1"), c)
	//helpers.GetRequest(fmt.Sprintf(env.DataServiceHost, "report1"), c)
}

func (h *Handler) report2(c *gin.Context) {
	helpers.GetRequest(fmt.Sprintf(env.DataServiceHost, "report2"), c)
}

func (h *Handler) report3(c *gin.Context) {
	helpers.GetRequest(fmt.Sprintf(env.DataServiceHost, "report3"), c)
}
