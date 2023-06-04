package usecase

import (
	"backend/model"
	"backend/repository"
)

type ILikeFlagUsecase interface {
	CreateLikeFlag(likeFlag model.Likeflag) (model.LikeflagResponse, error)
	AddLikeFlag(likeFlag model.Likeflag) (model.LikeflagResponse, error)
	AddUnlikeFlag(likeFlag model.Likeflag) (model.LikeflagResponse, error)
	GetIsLikedFlag(likeFlag model.Likeflag, account_id uint, track_id uint) (model.LikeflagResponse, error)
}

type likeFlagUsecase struct {
	lr repository.ILikeFlagRepository
}

func NewLikeFlagUsecase(atr repository.ILikeFlagRepository) ILikeFlagUsecase {
	return &likeFlagUsecase{atr}
}

func (lu likeFlagUsecase) CreateLikeFlag(likeFlag model.Likeflag) (model.LikeflagResponse, error) {
	if err := lu.lr.CreateLikeFlag(&likeFlag); err != nil {
		return model.LikeflagResponse{}, err
	}
	resAccountTrack := model.LikeflagResponse{
		ID:        likeFlag.ID,
		AccountID: likeFlag.AccountID,
		TrackID:   likeFlag.TrackID,
		Liked:     likeFlag.Liked,
	}
	return resAccountTrack, nil
}

func (lu likeFlagUsecase) AddLikeFlag(likeFlag model.Likeflag) (model.LikeflagResponse, error) {
	if err := lu.lr.AddLikeFlag(&likeFlag); err != nil {
		return model.LikeflagResponse{}, err
	}
	resAccountTrack := model.LikeflagResponse{
		ID:        likeFlag.ID,
		AccountID: likeFlag.AccountID,
		TrackID:   likeFlag.TrackID,
		Liked:     likeFlag.Liked,
	}
	return resAccountTrack, nil
}

func (lu likeFlagUsecase) AddUnlikeFlag(likeFlag model.Likeflag) (model.LikeflagResponse, error) {
	if err := lu.lr.AddUnlikeFlag(&likeFlag); err != nil {
		return model.LikeflagResponse{}, err
	}
	resAccountTrack := model.LikeflagResponse{
		ID:        likeFlag.ID,
		AccountID: likeFlag.AccountID,
		TrackID:   likeFlag.TrackID,
		Liked:     likeFlag.Liked,
	}
	return resAccountTrack, nil
}


func (lu likeFlagUsecase) GetIsLikedFlag(likeFlag model.Likeflag, account_id uint, track_id uint) (model.LikeflagResponse, error) {
	if err := lu.lr.GetIsLikedFlag(&likeFlag, account_id, track_id); err != nil {
		return model.LikeflagResponse{}, err
	}
	resAccountTrack := model.LikeflagResponse{
		ID:        likeFlag.ID,
		AccountID: likeFlag.AccountID,
		TrackID:   likeFlag.TrackID,
		Liked:     likeFlag.Liked,
	}
	return resAccountTrack, nil
}
