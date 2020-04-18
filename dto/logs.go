package dto

import "time"

type LogEntity struct {
	ID          uint      `gorm:"primary_key;auto_increment;" json:"id,omitempty"`
	Title       string    `gorm:"not null" json:"title,omitempty"`
	Description string    `gorm:"default:null" json:"description,omitempty"`
	Rating      int       `gorm:"not null; default:0" json:"rating,omitempty"`
	ImageUrl    string    `gorm:"not null" json:"image_url,omitempty"`
	Latitude    float64   `gorm:"not null; default:180" json:"latitude,omitempty"`
	Longitude   float64   `gorm:"not null; default:180" json:"longitude,omitempty"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}
