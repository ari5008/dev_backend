package controller

import (
	"backend/model"
	"backend/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ILikeFlagController interface {
	CreateLikeFlag(c echo.Context) error
	AddLikeFlag(c echo.Context) error
	AddUnLikeFlag(c echo.Context) error
	GetIsLikedFlag(c echo.Context) error
	DeleteLikeFlag(c echo.Context) error
}

type likeFlagController struct {
	lu usecase.ILikeFlagUsecase
	au usecase.IAccountUsecase
}

func NewLikeFlagController(lu usecase.ILikeFlagUsecase, au usecase.IAccountUsecase) ILikeFlagController {
	return &likeFlagController{lu, au}
}

func (lc likeFlagController) CreateLikeFlag(c echo.Context) error {
	likeFlag := model.Likeflag{}
	if err := c.Bind(&likeFlag); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	likeFlag.Liked = false
	resLikeFlag, err := lc.lu.CreateLikeFlag(likeFlag)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, resLikeFlag)
}

func (lc likeFlagController) AddLikeFlag(c echo.Context) error {
	likeFlag := model.Likeflag{}
	if err := c.Bind(&likeFlag); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resLikeFlag, err := lc.lu.AddLikeFlag(likeFlag)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, resLikeFlag)
	
}

func (lc likeFlagController) AddUnLikeFlag(c echo.Context) error {
	likeFlag := model.Likeflag{}
	if err := c.Bind(&likeFlag); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resLikeFlag, err := lc.lu.AddUnlikeFlag(likeFlag)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, resLikeFlag)
}

func (lc likeFlagController) GetIsLikedFlag(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	account, err := lc.au.GetAccount(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	id := c.Param("trackId")
	trackId, _ := strconv.Atoi(id)

	likeFlag := model.Likeflag{}
	if err := c.Bind(&likeFlag); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resLikeFlag, err := lc.lu.GetIsLikedFlag(likeFlag, uint(account.ID), uint(trackId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, resLikeFlag)
}

func (lc likeFlagController) DeleteLikeFlag(c echo.Context) error {

	id := c.Param("trackId")
	trackId, _ := strconv.Atoi(id)

	err := lc.lu.DeleteLikeFlag(uint(trackId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}