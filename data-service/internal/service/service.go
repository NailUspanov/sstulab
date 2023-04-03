package service

import (
	"data-service/internal/entity"
	"data-service/internal/storage"
)

type FilmServiceInterface interface {
	Create(film entity.Film) error
	QueueImplementation(msg any)
	GetTopFilmsByRating() ([]entity.FilmRating, error)
	GetTopDirectorsByFilms() ([]entity.DirectorFilms, error)
	GetTopFilmsByReviewsCount() ([]entity.FilmReviews, error)
}

type ReviewServiceInterface interface {
	Create(film entity.Review) error
}

type Service struct {
	FilmService   FilmServiceInterface
	ReviewService ReviewServiceInterface
}

func NewService(storage *storage.Storage) *Service {
	return &Service{FilmService: NewFilmService(storage.FilmsStorage), ReviewService: NewReviewService(storage.ReviewsStorage)}
}
