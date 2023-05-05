package router

import (
	"backend/controller"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController) *echo.Echo {
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
			CookiePath: "/",
			CookieDomain: os.Getenv("API_DOMAIN"),
			CookieHTTPOnly: true,
			// CookieSameSite: http.SameSiteNoneMode,
			CookieSameSite: http.SameSiteDefaultMode,
			// CookieMaxAge: 60,
		}))

	e.POST("/signup", uc.Signup)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.Logout)
	e.GET("/csrf", uc.CsrfToken)
	return e
}
