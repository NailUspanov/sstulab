package services

import (
	"github.com/Shopify/sarama"
)

type FilmServiceInterface interface {
	Create(film AddRequest) error
}

type ReviewServiceInterface interface {
	Create(film AddReviewRequest) error
}

type Service struct {
	FilmService   FilmServiceInterface
	ReviewService ReviewServiceInterface
}

func NewService(messageProducer sarama.SyncProducer) *Service {
	return &Service{
		FilmService:   NewFilmService(messageProducer),
		ReviewService: NewReviewService(messageProducer),
	}
}
