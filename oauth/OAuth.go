package oauth

import (
	"context"
	// "log"
	// "time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

var (
	token  *oauth2.Token
)

func AccessToken() (*oauth2.Token, error) {
	if token != nil && token.Valid() {
		return token, nil
	}

	ctx := context.Background()
	config := &clientcredentials.Config{
		ClientID:     "2ed6dfbde6e5476e886f4d4b328cba91",
		ClientSecret: "4b62455115c447caa0b7e8ddc3c7b42c",
		TokenURL:     "https://accounts.spotify.com/api/token",
		Scopes:       []string{"playlist-read-private", "playlist-read-collaborative", "streaming"},
	}
	token, err := config.Token(ctx)
	if err != nil {
		return nil, err
	}

	return token, nil
}

// func UpdateToken() {
// 	ticker := time.NewTicker(time.Minute) // 1時間ごとにトークンを更新
// 	defer ticker.Stop()
// 	for {
// 		<-ticker.C
// 		ctx := context.Background()
// 		var err error
// 		token, err = config.Token(ctx)
// 		if err != nil {
// 			log.Fatalf("Error while refreshing token: %v", err)
// 		}
// 	}
// }

