package repository

import (
	"backend/model"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITrackRepository interface {
	CreateTrack(track *model.Track) error
	GetAllTracks(tracks *[]model.Track) error
	GetTrackById(track *model.Track, trackId uint) error
	GetTrackByAccountId(track *model.Track, accountId uint) error
	UpdateTrack(track *model.Track, trackId uint) error
	IncrementSelectedTrackLikes(track *model.Track, trackId uint) error
	DecrementSelectedTrackLikes(track *model.Track, trackId uint) error
	DeleteTrack(accountId uint, trackId uint) error
}

type trackRepository struct {
	db *gorm.DB
}

func NewTrackRepository(db *gorm.DB) ITrackRepository {
	return &trackRepository{db}
}

func (tr *trackRepository) CreateTrack(track *model.Track) error {
	if err := tr.db.Create(track).Error; err != nil {
		return err
	}
	return nil
}

func (tr *trackRepository) GetAllTracks(tracks *[]model.Track) error {
	if err := tr.db.Order("created_at").Find(tracks).Error; err != nil {
		return err
	}
	return nil
}

func (tr *trackRepository) GetTrackById(track *model.Track, trackId uint) error {
	if err := tr.db.Where("id=?", trackId).First(track).Error; err != nil {
		return err
	}
	return nil
}

func (tr *trackRepository) GetTrackByAccountId(track *model.Track, accountId uint) error {
	if err := tr.db.Where("account_id=?", accountId).First(track).Error; err != nil {
		return err
	}
	return nil
}

func (tr *trackRepository) UpdateTrack(track *model.Track, trackId uint) error {
	result := tr.db.Model(track).Clauses(clause.Returning{}).Where("id=?", trackId).
		Updates(map[string]interface{}{
			"title":        track.Title,
			"artist_name":  track.ArtistName,
			"jacket_image": track.JacketImage,
			"genre":        track.Genre,
			"comment":      track.Comment,
			"likes":        track.Likes,
			"account_id":   track.AccountId,
		})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (tr *trackRepository) IncrementSelectedTrackLikes(track *model.Track, trackId uint) error {
	result := tr.db.Model(track).Clauses(clause.Returning{}).Where("id=?", trackId).
		Update("likes", int(track.Likes+1))
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (tr *trackRepository) DecrementSelectedTrackLikes(track *model.Track, trackId uint) error {
	result := tr.db.Model(track).Clauses(clause.Returning{}).Where("id=?", trackId).
		Update("likes", int(track.Likes-1))
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (tr *trackRepository) DeleteTrack(accountId uint, trackId uint) error {
	result := tr.db.Where("id=? AND account_id", trackId, accountId).Delete(&model.Track{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
