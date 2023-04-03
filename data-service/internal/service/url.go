package service

import (
	"context"
	"data-service/internal/entity"
	"data-service/internal/storage"
	"github.com/google/uuid"
)

type AddRequest struct {
	Name string `json:"name"`
}

type FilmService struct {
	storage storage.FilmsStorage
}

func NewFilmService(storage storage.FilmsStorage) *FilmService {
	return &FilmService{storage: storage}
}

func (u *FilmService) Create(film entity.Film) error {
	film.Id = uuid.New()
	err := u.storage.Add(context.TODO(), film)
	if err != nil {
		return err
	}
	return err
}

func (u *FilmService) QueueImplementation(msg any) {
}

func (u *FilmService) GetTopFilmsByRating() ([]entity.FilmRating, error) {
	return u.storage.GetTopFilmsByRating(context.TODO())
}

func (u *FilmService) GetTopDirectorsByFilms() ([]entity.DirectorFilms, error) {
	return u.storage.GetTopDirectorsByFilms(context.TODO())
}

func (u *FilmService) GetTopFilmsByReviewsCount() ([]entity.FilmReviews, error) {
	return u.storage.GetTopFilmsByReviewsCount(context.TODO())
}
