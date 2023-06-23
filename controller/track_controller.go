package controller

import (
	"backend/model"
	"backend/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ITrackController interface {
	CreateTrack(c echo.Context) error
	GetAllTracksByLikes(c echo.Context) error
	GetAllTracksByAsc(c echo.Context) error
	GetAllTracksByDesc(c echo.Context) error
	GetAllTracksByGenre(c echo.Context) error
	GetTrackByAccountId(c echo.Context) error
	UpdateTrack(c echo.Context) error
	DeleteTrack(c echo.Context) error
	IncrementSelectedTrackLikes(c echo.Context) error
	DecrementSelectedTrackLikes(c echo.Context) error
}

type trackController struct {
	tu usecase.ITrackUsecase
	au usecase.IAccountUsecase
}

func NewTrackController(tu usecase.ITrackUsecase, au usecase.IAccountUsecase) ITrackController {
	return &trackController{tu, au}
}

func (tc *trackController) CreateTrack(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	account, err := tc.au.GetAccount(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	track := model.Track{}
	if err := c.Bind(&track); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	
	track.AccountId = uint(account.ID)
	resTrack, err := tc.tu.CreateTrack(track)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, resTrack)
}

func (tc *trackController) GetAllTracksByLikes(c echo.Context) error {
	resTracks, err := tc.tu.GetAllTracksByLikes()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resTracks)
}

func (tc *trackController) GetAllTracksByAsc(c echo.Context) error {
	resTracks, err := tc.tu.GetAllTracksByAsc()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resTracks)
}

func (tc *trackController) GetAllTracksByDesc(c echo.Context) error {
	resTracks, err := tc.tu.GetAllTracksByDesc()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resTracks)
}

func (tc *trackController) GetAllTracksByGenre(c echo.Context) error {
	resTracks, err := tc.tu.GetAllTracksByGenre()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resTracks)
}

func (tc *trackController) GetTrackByAccountId(c echo.Context) error {
	id := c.Param("accountId")
	accountId, _ := strconv.Atoi(id)
	
	trackRes, err := tc.tu.GetTrackByAccountId(uint(accountId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, trackRes)
}

func (tc *trackController) UpdateTrack(c echo.Context) error {
	id := c.Param("trackId")
	trackId, _ := strconv.Atoi(id)
	
	track := model.Track{}
	if err := c.Bind(&track); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	trackRes, err := tc.tu.UpdateTrack(track, uint(trackId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, trackRes)
}

func (tc *trackController) DeleteTrack(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	account, err := tc.au.GetAccount(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	id := c.Param("trackId")
	trackId, _ := strconv.Atoi(id)
	
	err = tc.tu.DeleteTrack(uint(account.ID), uint(trackId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}



func (tc *trackController) IncrementSelectedTrackLikes(c echo.Context) error {
	id := c.Param("trackId")
	trackId, _ := strconv.Atoi(id)
	
	track := model.Track{}
	if err := c.Bind(&track); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	trackRes, err := tc.tu.IncrementSelectedTrackLikes(track, uint(trackId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, trackRes)
}

func (tc *trackController) DecrementSelectedTrackLikes(c echo.Context) error {
	id := c.Param("trackId")
	trackId, _ := strconv.Atoi(id)
	
	track := model.Track{}
	if err := c.Bind(&track); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	trackRes, err := tc.tu.DecrementSelectedTrackLikes(track, uint(trackId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, trackRes)
}
