package postgres

import (
	"context"
	"data-service/internal/entity"
	"errors"
	"gorm.io/gorm"
)

type ReviewsStorage struct {
	db *gorm.DB
}

func NewReviewsStorage(db *gorm.DB) *ReviewsStorage {
	return &ReviewsStorage{db: db}
}

func (r *ReviewsStorage) Add(ctx context.Context, review entity.Review) error {
	return r.db.WithContext(ctx).Table(entity.Review{}.TableName()).Create(&review).Error
}

func (r *ReviewsStorage) GetAll(ctx context.Context) ([]entity.Review, error) {
	var res []entity.Review
	db := r.db.WithContext(ctx).Table(entity.Review{}.TableName()).Select("*").Find(&res)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, db.Error
		}
	}
	return res, nil
}
