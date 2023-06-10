package router

import (
	"backend/controller"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, ac controller.IAccountController, tc controller.ITrackController, lc controller.ILikeFlagController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http:localhost:3000",
			os.Getenv("Front_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.
		CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		// CookieSameSite: http.SameSiteNoneMode,
		CookieSameSite: http.SameSiteDefaultMode,
		// CookieMaxAge: 60,
	}))

	e.POST("/signup", uc.Signup)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.Logout)
	e.GET("/csrf", uc.CsrfToken)

	e.GET("searchApi", controller.GetSearchResults)

	a := e.Group("/account")
	a.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	
	a.GET("", ac.GetAccount)	
	a.PUT("/:accountId", ac.UpdateAccount)
	a.DELETE("/:accountId", ac.DeleteAccount)
	
	a.GET("/tracks", tc.GetAllTracks)
	a.GET("/track/:trackId", tc.GetTrackById)
	a.POST("/createTrack", tc.CreateTrack)
	a.PUT("/updateTrack/:trackId", tc.UpdateTrack)
	a.PUT("/incrementTrackLikes/:trackId", tc.IncrementSelectedTrackLikes)
	a.PUT("/decrementTrackLikes/:trackId", tc.DecrementSelectedTrackLikes)

	a.POST("/createLikeFlag", lc.CreateLikeFlag)
	a.PUT("/addLikeFlag", lc.AddLikeFlag)
	a.PUT("/addUnLikeFlag", lc.AddUnLikeFlag)
	a.GET("/getLikeFlag/:trackId", lc.GetIsLikedFlag)

	return e
}
