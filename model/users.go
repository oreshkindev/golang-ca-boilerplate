package model

import "time"

type Users struct {
	ID        int64     `json:"id"`
	Email     string    `gorm:"uniqueIndex:idx_user_email,sort:desc" json:"email"`
	Password  string    `json:"password"`
	Status    bool      `gorm:"default:false" json:"Status"`
	CreatedAt time.Time `gorm:"default:current_timestamp" json:"createdAt"`
	UpdatedAt time.Time `gorm:"default:current_timestamp" json:"updatedAt"`
}
