package oauth

import (
	"context"
	"os"
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
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		TokenURL:     os.Getenv("TokenURL"),
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

