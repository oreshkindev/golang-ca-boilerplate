package model

import (
	"time"
)

type Posts struct {
	ID          int32     `json:"id"`
	Title       string    `gorm:"uniqueIndex:idx_post_title,sort:desc" json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Author      string    `json:"author"`
	Image       string    `json:"image"`
	Status      bool      `gorm:"default:false" json:"status"`
	CreatedAt   time.Time `gorm:"default:current_timestamp" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"default:current_timestamp" json:"updatedAt"`
}
