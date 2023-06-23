package controller

import (
	"backend/model"
	"backend/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IAccountController interface {
	GetAccount(c echo.Context) error
	GetAccountByTrackId(c echo.Context) error
	UpdateAccount(c echo.Context) error
	DeleteAccount(c echo.Context) error
}

type accountController struct {
	au usecase.IAccountUsecase
	tu usecase.ITrackUsecase
}

func NewAccountController(au usecase.IAccountUsecase, tu usecase.ITrackUsecase) IAccountController {
	return &accountController{au, tu}
}

func (ac *accountController) GetAccount(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	resAccount, err := ac.au.GetAccount(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resAccount)
}

func (ac *accountController) GetAccountByTrackId(c echo.Context) error {
	id := c.Param("trackId")
	trackId, _ := strconv.Atoi(id)
	resTrack, err := ac.tu.GetTrackById(uint(trackId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	resAccount, err := ac.au.GetAccountById(uint(resTrack.AccountId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resAccount)
}

func (ac *accountController) UpdateAccount(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("accountId")
	accountId, _ := strconv.Atoi(id)

	account := model.Account{}
	if err := c.Bind(&account); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	accountRes, err := ac.au.UpdateAccount(account, uint(userId.(float64)), uint(accountId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, accountRes)
}

func (ac *accountController) DeleteAccount(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("accountId")
	accountId, _ := strconv.Atoi(id)

	err := ac.au.DeleteAccount(uint(userId.(float64)), uint(accountId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
