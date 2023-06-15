package controller

import (
	"backend/oauth"
	"backend/utils"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type search struct {
	Tracks struct {
		Items []struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			External_urls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Artists []struct {
				Name string `json:"name"`
			} `json:"artists"`
			Album struct {
				Images []struct {
					URL string `json:"url"`
				} `json:"images"`
			} `json:"album"`
		} `json:"items"`
	} `json:"tracks"`
}

type AccountTrack struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	External_url string `json:"external_url"`
	Artists      string `json:"artists"`
	ImageURL     string `json:"image_url"`
}

// func GetToken(c echo.Context) error {

// 	token, err := oauth.AccessToken()
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}

// 	jwtToken := jwt.New(jwt.SigningMethodHS256)
// 	claims := jwtToken.Claims.(jwt.MapClaims)
// 	claims["sub"] = "access_token"
// 	claims["exp"] = token.Expiry.Unix()
// 	claims["access_token"] = token.AccessToken
// 	claims["token_type"] = token.TokenType

// 	secretKey := []byte(os.Getenv("SECRET2"))
// 	tokenString, err := jwtToken.SignedString(secretKey)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}

// 	originalToken, err := utils.ValidateToken(tokenString, secretKey)
// 	if err != nil {
// 		return c.JSON(http.StatusUnauthorized, err.Error())
// 	}

// 	accessToken := originalToken.Claims.(jwt.MapClaims)["access_token"].(string)
// 	return c.JSON(http.StatusOK, accessToken)
// }


func GetSearchResults(c echo.Context) error {

	q := c.QueryParam("q")

	token, err := oauth.AccessToken()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	client := &http.Client{}
	baseURL := "https://api.spotify.com/v1/search"
	fullURL := utils.ParamsUrl(q, baseURL)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	req.Header.Set("Accept-Language", "ja")
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	resp, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer resp.Body.Close()

	var s search
	err = json.NewDecoder(resp.Body).Decode(&s)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	accountTracks := []AccountTrack{}
	for _, t := range s.Tracks.Items {
		artistNames := make([]string, len(t.Artists))
		for i, a := range t.Artists {
			artistNames[i] = a.Name
		}
		artist := strings.Join(artistNames, ",")
		track := AccountTrack{
			ID:           t.ID,
			Name:         t.Name,
			External_url: t.External_urls.Spotify,
			Artists:      artist,
			ImageURL:     "",
		}
		if len(t.Album.Images) > 0 {
			track.ImageURL = t.Album.Images[1].URL
		}
		accountTracks = append(accountTracks, track)
	}

	return c.JSON(http.StatusOK, accountTracks)
}
