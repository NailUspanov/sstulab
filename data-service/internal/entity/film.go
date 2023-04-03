package entity

import "github.com/google/uuid"

type Film struct {
	Id       uuid.UUID `json:"id" gorm:"primaryKey; type:uuid; unique; not null"`
	Name     string    `json:"name" gorm:"not null"`
	Director string    `json:"director"`
}

func (f Film) TableName() string {
	return "films"
}

type FilmRating struct {
	Film
	AvgRating float64 `json:"avg_rating"`
}

type DirectorFilms struct {
	Director string `json:"director"`
	NumFilms int    `json:"num_films"`
}

type FilmReviews struct {
	Name       string `json:"name"`
	NumReviews int    `json:"num_reviews"`
}
