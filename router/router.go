package router

import (
	"backend/controller"
	"backend/utils"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, ac controller.IAccountController, tc controller.ITrackController, lc controller.ILikeFlagController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000",
			os.Getenv("Front_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken, "Authorization"},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userAgent := c.Request().Header.Get("User-Agent")
			isMobile := utils.IsMobileDevice(userAgent)
			if isMobile {
				return next(c)
			} else {
				return middleware.CSRFWithConfig(middleware.CSRFConfig{
					CookiePath:     "/",
					CookieDomain:   os.Getenv("API_DOMAIN"),
					CookieHTTPOnly: true,
					CookieSameSite: http.SameSiteNoneMode,
					// CookieSameSite: http.SameSiteDefaultMode,
					CookieMaxAge: 300,
				})(next)(c)
			}
		}
	})

	e.POST("/signup", uc.Signup)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.Logout)
	e.GET("/csrf", uc.CsrfToken)

	e.GET("/tracks", tc.GetAllTracks)
	e.GET("/getAccount/:trackId", ac.GetAccountByTrackId)
	e.GET("searchApi", controller.GetSearchResults)

	a := e.Group("/account")
	a.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))

	a.GET("", ac.GetAccount)
	a.PUT("/:accountId", ac.UpdateAccount)
	a.DELETE("/:accountId", ac.DeleteAccount)

	a.GET("/trackByAccountId/:accountId", tc.GetTrackByAccountId)
	a.POST("/createTrack", tc.CreateTrack)
	a.DELETE("/deleteTrack/:trackId", tc.DeleteTrack)
	a.PUT("/incrementTrackLikes/:trackId", tc.IncrementSelectedTrackLikes)
	a.PUT("/decrementTrackLikes/:trackId", tc.DecrementSelectedTrackLikes)

	a.POST("/createLikeFlag", lc.CreateLikeFlag)
	a.PUT("/addLikeFlag", lc.AddLikeFlag)
	a.PUT("/addUnLikeFlag", lc.AddUnLikeFlag)
	a.GET("/getLikeFlag/:trackId", lc.GetIsLikedFlag)
	a.DELETE("/deleteLikeFlag/:trackId", lc.DeleteLikeFlag)

	return e
}
