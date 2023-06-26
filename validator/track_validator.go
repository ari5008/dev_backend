package validator

import (
	"backend/model"
	"backend/repository"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ITrackValidator interface {
	TrackValidate(track model.Track) error
}

type trackValidator struct{
	tr repository.ITrackRepository
}

func NewTrackValidator(tr repository.ITrackRepository) ITrackValidator {
	return &trackValidator{tr}
}

func (tv *trackValidator) TrackValidate(track model.Track) error {
	err := tv.tr.NotSameTitleAndAccountID(&track)
	if err != nil {
		return err
	}
	err = tv.tr.NotSameTrack(&track)
	if err != nil {
		return err
	}
	return validation.ValidateStruct(&track,
		validation.Field(
			&track.Title,
			validation.Required.Error("require title"),
		),
		validation.Field(
			&track.ArtistName,
			validation.Required.Error("require artist_name"),
		),
		validation.Field(
			&track.JacketImage,
			validation.Required.Error("require jacket_image"),
		),
		validation.Field(
			&track.Genre,
			validation.Required.Error("require genre"),
		),
		validation.Field(
			&track.Comment,
			validation.RuneLength(0, 100).Error("limit comment"),
		),
	)
}