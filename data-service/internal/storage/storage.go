package storage

import (
	"context"
	"data-service/internal/entity"
	"data-service/internal/storage/postgres"
	"gorm.io/gorm"
)

type FilmsStorage interface {
	Add(ctx context.Context, film entity.Film) error
	GetAll(ctx context.Context) ([]entity.Film, error)
	GetTopFilmsByRating(ctx context.Context) ([]entity.FilmRating, error)
	GetTopDirectorsByFilms(ctx context.Context) ([]entity.DirectorFilms, error)
	GetTopFilmsByReviewsCount(ctx context.Context) ([]entity.FilmReviews, error)
}

type ReviewsStorage interface {
	Add(ctx context.Context, review entity.Review) error
	GetAll(ctx context.Context) ([]entity.Review, error)
}

type Storage struct {
	ReviewsStorage
	FilmsStorage
}

func NewStorage(db *gorm.DB) *Storage {
	return &Storage{FilmsStorage: postgres.NewFilmsStorage(db), ReviewsStorage: postgres.NewReviewsStorage(db)}
}
