package model

import (
	"time"
)

type Likeflag struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	AccountID uint      `json:"account_id" gorm:"index"`
	TrackID   uint      `json:"track_id" gorm:"index"`
	Liked     bool      `json:"liked" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type LikeflagResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	AccountID uint      `json:"account_id" gorm:"index"`
	TrackID   uint      `json:"track_id" gorm:"index"`
	Liked     bool      `json:"liked" gorm:"default:false"`
}
