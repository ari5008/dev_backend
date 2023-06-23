package model

import "time"

type Track struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Title        string    `json:"title" gorm:"not null"`
	ArtistName   string    `json:"artist_name"`
	JacketImage  string    `json:"jacket_image"`
	Genre        string    `json:"genre"`
	Comment      string    `json:"comment"`
	Likes        int       `json:"likes"`
	External_url string    `json:"external_url"`
	AccountId    uint      `json:"account_id" gorm:"not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type TrackResponse struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Title        string    `json:"title" gorm:"not null"`
	ArtistName   string    `json:"artist_name"`
	JacketImage  string    `json:"jacket_image"`
	Genre        string    `json:"genre"`
	Comment      string    `json:"comment"`
	Likes        int       `json:"likes"`
	External_url string    `json:"external_url"`
	AccountId    uint      `json:"account_id" gorm:"not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
