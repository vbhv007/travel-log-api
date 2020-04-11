package dto

import "time"

type LogEntity struct {
	ID          uint      `gorm:"primary_key;auto_increment;" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `gorm:"-" json:"description"`
	Rating      int       `gorm:"not null; default:0" json:"latitude"`
	ImageUrl    string    `gorm:"not null" json:"image_url"`
	Latitude    int       `gorm:"not null; default:180" json:"latitude"`
	Longitude   int       `gorm:"not null; default:180" json:"longitude"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
