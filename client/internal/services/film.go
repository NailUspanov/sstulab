package services

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"net/http"
)

type AddRequest struct {
	RoutingKey string  `json:"routingKey"`
	Text       string  `json:"text"`
	Rating     float64 `json:"rating"`
	FilmId     string  `json:"filmId"`
}

type FilmService struct {
	messageProducer sarama.SyncProducer
	client          *http.Client
}

func NewFilmService(messageProducer sarama.SyncProducer) *FilmService {
	return &FilmService{messageProducer: messageProducer}
}

func (u *FilmService) Create(film AddRequest) error {
	jsonMessage, err := json.Marshal(film)
	if err != nil {
		log.Fatal(err)
	}
	msg := &sarama.ProducerMessage{
		Topic: "asd",
		Value: sarama.StringEncoder(jsonMessage),
	}
	partition, offset, err := u.messageProducer.SendMessage(msg)
	if err != nil {
		log.Fatalf("Error sending message to Kafka: %s", err.Error())
	}
	fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)
	return nil
}
