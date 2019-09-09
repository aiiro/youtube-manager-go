package models

import (
	"time"
)

type Favorite struct {
	ID        uint      `gorm:"primary_key"`
	UserId    uint      `json:"user_id"`
	VideoId   string    `json:"video_id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	User User
}
