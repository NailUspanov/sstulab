package entity

import "github.com/google/uuid"

type Review struct {
	Id     uuid.UUID  `json:"id" gorm:"primaryKey; type:uuid; unique; not null"`
	Text   string     `json:"text" gorm:"not null"`
	Rating float64    `json:"rating" gorm:"not null; index"`
	FilmId *uuid.UUID `json:"filmId" gorm:"not null; index; type:uuid"`
	Film   *Film      `json:"film" gorm:"constraint:OnDelete:CASCADE;"`
}

func (r Review) TableName() string {
	return "reviews"
}
