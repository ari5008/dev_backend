package repository

import (
	"backend/model"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITrackRepository interface {
	CreateTrack(track *model.Track) error
	GetAllTracksByLikes(tracks *[]model.Track) error
	GetAllTracksByAsc(tracks *[]model.Track) error
	GetAllTracksByDesc(tracks *[]model.Track) error
	GetAllTracksByGenre(tracks *[]model.Track) error
	GetTrackById(track *model.Track, trackId uint) error
	GetTrackByAccountId(tracks *[]model.Track, accountId uint) error
	UpdateTrack(track *model.Track, trackId uint) error
	DeleteTrack(accountId uint, trackId uint) error
	IncrementSelectedTrackLikes(track *model.Track, trackId uint) error
	DecrementSelectedTrackLikes(track *model.Track, trackId uint) error
	NotSameTitleAndAccountID(track *model.Track) error
	NotSameTrack(track *model.Track) error
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

func (tr *trackRepository) GetAllTracksByLikes(tracks *[]model.Track) error {
	if err := tr.db.Order("likes DESC").Find(tracks).Error; err != nil {
		return err
	}
	return nil
}

func (tr *trackRepository) GetAllTracksByAsc(tracks *[]model.Track) error {
	if err := tr.db.Order("created_at ASC").Find(tracks).Error; err != nil {
		return err
	}
	return nil
}

func (tr *trackRepository) GetAllTracksByDesc(tracks *[]model.Track) error {
	if err := tr.db.Order("created_at DESC").Find(tracks).Error; err != nil {
		return err
	}
	return nil
}

func (tr *trackRepository) GetAllTracksByGenre(tracks *[]model.Track) error {
	if err := tr.db.Order("CASE WHEN genre = 'ポップ' THEN 1 WHEN genre = '可愛い' THEN 2 WHEN genre = 'ロック' THEN 3 WHEN genre = 'ヒップホップ' THEN 4 WHEN genre = 'レトロ' THEN 5 WHEN genre = 'アンニュイ' THEN 6 WHEN genre = '癒されたい' THEN 7 WHEN genre = 'テンションが上がる' THEN 8 WHEN genre = '無心で聞きたい' THEN 9 WHEN genre = 'ドライブで聞きたい' THEN 10 WHEN genre = '最近のおすすめ' THEN 11 END").Find(tracks).Error; err != nil {
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

func (tr *trackRepository) GetTrackByAccountId(tracks *[]model.Track, accountId uint) error {
	if err := tr.db.Order("created_at").Where("account_id=?", accountId).Find(tracks).Error; err != nil {
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
			"external_url": track.External_url,
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

func (tr *trackRepository) DeleteTrack(accountId uint, trackId uint) error {
	result := tr.db.Where("id=? AND account_id=?", trackId, accountId).Delete(&model.Track{})
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

func (tr *trackRepository) NotSameTitleAndAccountID(track *model.Track) error {
	var count int64
	if err := tr.db.Model(&model.Track{}).Where("title = ? AND artist_name = ? AND account_id=?", track.Title, track.ArtistName, track.AccountId).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		// 同じ曲があったらエラーを返す
		return fmt.Errorf("duplicated Title and ArtistName and AccountID")
	}
	return nil

}
func (tr *trackRepository) NotSameTrack(track *model.Track) error {
	var count int64
	if err := tr.db.Model(&model.Track{}).Where("title = ? AND artist_name = ? AND genre = ?", track.Title, track.ArtistName, track.Genre).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		// 同じ曲があったらエラーを返す
		return fmt.Errorf("duplicated track")
	}
	return nil

}
