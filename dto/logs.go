package dto

import "time"

type LogEntity struct {
	ID          uint      `gorm:"primary_key;auto_increment;" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `gorm:"-" json:"description"`
	Rating      int       `gorm:"not null; default:0" json:"latitude"`
	ImageUrl    string    `gorm:"not null" json:"image_url"`
	Latitude    float64   `gorm:"not null; default:1000000000" json:"latitude"`
	Longitude   float64   `gorm:"not null; default:1000000000" json:"longitude"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
