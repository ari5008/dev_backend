package usecase

import (
	"backend/model"
	"backend/repository"
	"backend/validator"
)

type ITrackUsecase interface {
	CreateTrack(track model.Track) (model.TrackResponse, error)
	GetAllTracks() ([]model.Track, error)
	GetTrackById(trackId uint) (model.TrackResponse, error)
	GetTrackByAccountId(accountId uint) ([]model.Track, error)
	DeleteTrack(accountId uint, trackId uint) error
	IncrementSelectedTrackLikes(track model.Track, trackId uint) (model.TrackResponse, error)
	DecrementSelectedTrackLikes(track model.Track, trackId uint) (model.TrackResponse, error)
}

type trackUsecase struct {
	tr repository.ITrackRepository
	tv validator.ITrackValidator
}

func NewTrackUsecase(tr repository.ITrackRepository, tv validator.ITrackValidator) ITrackUsecase {
	return &trackUsecase{tr, tv}
}

func (tu *trackUsecase) CreateTrack(track model.Track) (model.TrackResponse, error) {
	if err := tu.tv.TrackValidate(track); err != nil {
		return model.TrackResponse{}, err
	}
	if err := tu.tr.CreateTrack(&track); err != nil {
		return model.TrackResponse{}, err
	}
	resTrack := model.TrackResponse{
		ID:           track.ID,
		Title:        track.Title,
		ArtistName:   track.ArtistName,
		JacketImage:  track.JacketImage,
		Genre:        track.Genre,
		Comment:      track.Comment,
		Likes:        track.Likes,
		External_url: track.External_url,
		AccountId:    track.AccountId,
		CreatedAt:    track.CreatedAt,
	}
	return resTrack, nil
}

func (tu *trackUsecase) GetAllTracks() ([]model.Track, error) {
	tracks := []model.Track{}
	if err := tu.tr.GetAllTracks(&tracks); err != nil {
		return nil, err
	}
	resTracks := []model.Track{}
	for _, v := range tracks {
		t := model.Track{
			ID:           v.ID,
			Title:        v.Title,
			ArtistName:   v.ArtistName,
			JacketImage:  v.JacketImage,
			Genre:        v.Genre,
			Comment:      v.Comment,
			Likes:        v.Likes,
			External_url: v.External_url,
			AccountId:    v.AccountId,
			CreatedAt:    v.CreatedAt,
		}
		resTracks = append(resTracks, t)
	}
	return resTracks, nil
}

func (tu *trackUsecase) GetTrackById(trackId uint) (model.TrackResponse, error) {
	track := model.Track{}
	if err := tu.tr.GetTrackById(&track, trackId); err != nil {
		return model.TrackResponse{}, err
	}
	resTrack := model.TrackResponse{
		ID:           track.ID,
		Title:        track.Title,
		ArtistName:   track.ArtistName,
		JacketImage:  track.JacketImage,
		Genre:        track.Genre,
		Comment:      track.Comment,
		Likes:        track.Likes,
		External_url: track.External_url,
		AccountId:    track.AccountId,
		CreatedAt:    track.CreatedAt,
	}
	return resTrack, nil
}

func (tu *trackUsecase) GetTrackByAccountId(accountId uint) ([]model.Track, error) {
	tracks := []model.Track{}
	if err := tu.tr.GetTrackByAccountId(&tracks, accountId); err != nil {
		return nil, err
	}
	resTracks := []model.Track{}
	for _, v := range tracks {
		t := model.Track{
			ID:           v.ID,
			Title:        v.Title,
			ArtistName:   v.ArtistName,
			JacketImage:  v.JacketImage,
			Genre:        v.Genre,
			Comment:      v.Comment,
			Likes:        v.Likes,
			External_url: v.External_url,
			AccountId:    v.AccountId,
			CreatedAt:    v.CreatedAt,
		}
		resTracks = append(resTracks, t)
	}
	return resTracks, nil
}

func (tu *trackUsecase) DeleteTrack(accountId uint, trackId uint) error {
	if err := tu.tr.DeleteTrack(accountId, trackId); err != nil {
		return err
	}
	return nil
}

func (tu *trackUsecase) IncrementSelectedTrackLikes(track model.Track, trackId uint) (model.TrackResponse, error) {
	if err := tu.tr.IncrementSelectedTrackLikes(&track, trackId); err != nil {
		return model.TrackResponse{}, err
	}
	resTrack := model.TrackResponse{
		ID:           track.ID,
		Title:        track.Title,
		ArtistName:   track.ArtistName,
		JacketImage:  track.JacketImage,
		Genre:        track.Genre,
		Comment:      track.Comment,
		Likes:        track.Likes,
		External_url: track.External_url,
		AccountId:    track.AccountId,
		CreatedAt:    track.CreatedAt,
	}
	return resTrack, nil
}

func (tu *trackUsecase) DecrementSelectedTrackLikes(track model.Track, trackId uint) (model.TrackResponse, error) {
	if err := tu.tr.DecrementSelectedTrackLikes(&track, trackId); err != nil {
		return model.TrackResponse{}, err
	}
	resTrack := model.TrackResponse{
		ID:           track.ID,
		Title:        track.Title,
		ArtistName:   track.ArtistName,
		JacketImage:  track.JacketImage,
		Genre:        track.Genre,
		Comment:      track.Comment,
		Likes:        track.Likes,
		External_url: track.External_url,
		AccountId:    track.AccountId,
		CreatedAt:    track.CreatedAt,
	}
	return resTrack, nil
}
