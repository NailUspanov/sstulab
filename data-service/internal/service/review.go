package service

import (
	"context"
	"data-service/internal/entity"
	"data-service/internal/storage"
	"github.com/google/uuid"
)

type AddReviewRequest struct {
	Rating float64 `json:"rating"`
	Text   string  `json:"text"`
}

type ReviewService struct {
	storage storage.ReviewsStorage
}

func NewReviewService(storage storage.ReviewsStorage) *ReviewService {
	return &ReviewService{storage: storage}
}

func (u *ReviewService) Create(review entity.Review) error {
	review.Id = uuid.New()
	return u.storage.Add(context.TODO(), review)
}
