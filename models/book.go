package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint   ` json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	MediumPrice string `json:"medium_price"`
	Author      string `json:"author"`
	ImageURL    string `json:"image_url"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
