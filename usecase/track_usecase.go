package usecase

import (
	"backend/model"
	"backend/repository"
)

type ITrackUsecase interface {
	CreateTrack(track model.Track) (model.TrackResponse, error)
	GetAllTracks() ([]model.TrackResponse, error)
	GetTrackById(accountId uint) (model.TrackResponse, error)
	UpdateTrack(track model.Track, accountId uint, trackId uint) (model.TrackResponse, error)
	DeleteTrack(accountId uint, trackId uint) error
}

type trackUsecase struct {
	tr repository.ITrackRepository
}

func NewTrackUsecase(tr repository.ITrackRepository) ITrackUsecase {
	return &trackUsecase{tr}
}

func (tu *trackUsecase) CreateTrack(track model.Track) (model.TrackResponse, error) {
	if err := tu.tr.CreateTrack(&track); err != nil {
		return model.TrackResponse{}, err
	}
	resTrack := model.TrackResponse{
		ID:          track.ID,
		Title:       track.Title,
		ArtistName:  track.ArtistName,
		JacketImage: track.JacketImage,
		Genre:       track.Genre,
		Likes:       track.Likes,
	}
	return resTrack, nil
}

func (tu *trackUsecase) GetAllTracks() ([]model.TrackResponse, error) {
	tracks := []model.Track{}
	if err := tu.tr.GetAllTracks(&tracks); err != nil {
		return nil, err
	}
	resTracks := []model.TrackResponse{}
	for _, v := range tracks {
		t := model.TrackResponse{
			ID:          v.ID,
			Title:       v.Title,
			ArtistName:  v.ArtistName,
			JacketImage: v.JacketImage,
			Genre:       v.Genre,
			Likes:       v.Likes,
		}
		resTracks = append(resTracks, t)
	}
	return resTracks, nil
}

func (tu *trackUsecase) GetTrackById(accountId uint) (model.TrackResponse, error) {
	track := model.Track{}
	if err := tu.tr.GetTrackById(&track, accountId); err != nil {
		return model.TrackResponse{}, err
	}
	resTrack := model.TrackResponse{
		ID:          track.ID,
		Title:       track.Title,
		ArtistName:  track.ArtistName,
		JacketImage: track.JacketImage,
		Genre:       track.Genre,
		Likes:       track.Likes,
	}
	return resTrack, nil
}

func (tu *trackUsecase) UpdateTrack(track model.Track, accountId uint, trackId uint) (model.TrackResponse, error) {
	if err := tu.tr.UpdateTrack(&track, accountId, trackId); err != nil {
		return model.TrackResponse{}, err
	}
	resTrack := model.TrackResponse{
		ID:          track.ID,
		Title:       track.Title,
		ArtistName:  track.ArtistName,
		JacketImage: track.JacketImage,
		Genre:       track.Genre,
		Likes:       track.Likes,
	}
	return resTrack, nil
}

func (tu *trackUsecase) DeleteTrack(accountId uint, trackId uint) error {
	if err := tu.tr.DeleteTrack(accountId, trackId); err != nil {
		return err
	}
	return nil
}
