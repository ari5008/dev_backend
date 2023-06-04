package repository

import (
	"backend/model"

	"gorm.io/gorm"
)

type ILikeFlagRepository interface {
	CreateLikeFlag(likeFlag *model.Likeflag) error
	AddLikeFlag(likeFlag *model.Likeflag) error
	AddUnlikeFlag(likeFlag *model.Likeflag) error
	GetIsLikedFlag(likeFlag *model.Likeflag, account_id uint, track_id uint) error
}

type likeFlagRepository struct {
	db *gorm.DB
}

func NewLikeFlagRepository(db *gorm.DB) ILikeFlagRepository {
	return &likeFlagRepository{db}
}

func (lr *likeFlagRepository) CreateLikeFlag(likeFlag *model.Likeflag) error {

	var count int64
	result := lr.db.Model(likeFlag).Where("account_id = ? AND track_id = ?", likeFlag.AccountID, likeFlag.TrackID).Count(&count)
	if result.Error != nil {
		return nil
	}

	if count == 0 {
		if err := lr.db.Create(&likeFlag).Error; err != nil {
			return err
		}
	}
	return nil
}

func (lr *likeFlagRepository) AddLikeFlag(likeFlag *model.Likeflag) error {
	if err := lr.db.Model(likeFlag).
		Where("account_id = ? AND track_id = ?", likeFlag.AccountID, likeFlag.TrackID).
		Update("liked", true).Error; err != nil {
		return err
	}
	return nil
}

func (lr *likeFlagRepository) AddUnlikeFlag(likeFlag *model.Likeflag) error {
	if err := lr.db.Model(likeFlag).
		Where("account_id = ? AND track_id = ?", likeFlag.AccountID, likeFlag.TrackID).
		Update("liked", false).Error; err != nil {
		return err
	}
	return nil
}

func (lr *likeFlagRepository) GetIsLikedFlag(likeFlag *model.Likeflag, account_id uint, track_id uint) error {
	if err := lr.db.Where("account_id = ? AND track_id = ?", account_id, track_id).First(likeFlag).Error; err != nil {
		return err
	}
	return nil
}
