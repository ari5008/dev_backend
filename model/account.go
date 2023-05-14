package model

import "time"

type Account struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	ImageURL     string    `json:"image_url"`
	Introduction string    `json:"introduction"`
	UserId       uint      `json:"user_id"  gorm:"not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type AccountResponse struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	ImageURL     string    `json:"image_url"`
	Introduction string    `json:"introduction"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
