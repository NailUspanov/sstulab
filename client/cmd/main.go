package main

import (
	"client/internal/handlers"
	"client/internal/helpers"
	"client/internal/services"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{"broker:9092"}, config)
	if err != nil {
		log.Fatalf("Error creating Kafka producer: %s", err.Error())
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalf("Error closing Kafka producer: %s", err.Error())
		}
	}()

	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	fmt.Println("HERE1")
	services := services.NewService(producer)
	handlers := handlers.NewHandler(services, &helpers.CircuitBreaker{
		MaxErrors:   3,
		Timeout:     10 * time.Second,
		ResetTime:   30 * time.Second,
		Status:      "closed",
		ErrorCount:  0,
		LastFailure: time.Time{},
	})
	fmt.Println("HERE")
	srv := new(Server)
	if err := srv.Run("8089", handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while runnning http server: %s", err.Error())
	}
}
