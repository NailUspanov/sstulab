package postgres

import (
	"context"
	"data-service/internal/entity"
	"errors"
	"gorm.io/gorm"
)

type FilmsStorage struct {
	db *gorm.DB
}

func NewFilmsStorage(db *gorm.DB) *FilmsStorage {
	return &FilmsStorage{db: db}
}

func (f *FilmsStorage) Add(ctx context.Context, film entity.Film) error {
	return f.db.WithContext(ctx).Table(entity.Film{}.TableName()).Create(&film).Error
}

func (f *FilmsStorage) GetAll(ctx context.Context) ([]entity.Film, error) {
	var res []entity.Film
	db := f.db.WithContext(ctx).Table(entity.Film{}.TableName()).Select("*").Find(&res)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, db.Error
		}
	}
	return res, nil
}

func (f *FilmsStorage) GetTopFilmsByRating(ctx context.Context) ([]entity.FilmRating, error) {
	var res []entity.FilmRating
	sql := `
		SELECT films.name, AVG(reviews.rating) AS avg_rating
		FROM films
		JOIN reviews ON films.id = reviews.film_id
		GROUP BY films.name
		ORDER BY avg_rating DESC
		LIMIT 10;
	`
	db := f.db.WithContext(ctx).Raw(sql).Find(&res)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, db.Error
		}
	}
	return res, nil
}

func (f *FilmsStorage) GetTopDirectorsByFilms(ctx context.Context) ([]entity.DirectorFilms, error) {
	var res []entity.DirectorFilms
	sql := `
		SELECT director, COUNT(*) AS num_films
		FROM films
		GROUP BY director
		ORDER BY num_films DESC
		LIMIT 10;
	`
	db := f.db.WithContext(ctx).Raw(sql).Find(&res)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, db.Error
		}
	}
	return res, nil
}

func (f *FilmsStorage) GetTopFilmsByReviewsCount(ctx context.Context) ([]entity.FilmReviews, error) {
	var res []entity.FilmReviews
	sql := `
		SELECT films.name, COUNT(reviews.id) AS num_reviews
		FROM films
		JOIN reviews ON films.id = reviews.film_id
		GROUP BY films.name
		ORDER BY num_reviews DESC
		LIMIT 10;
	`
	db := f.db.WithContext(ctx).Raw(sql).Find(&res)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, db.Error
		}
	}
	return res, nil
}
