package oauth

import (
	"context"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

var (
	token *oauth2.Token
)

func AccessToken() (*oauth2.Token, error) {
	if token != nil && !token.Expiry.IsZero() && token.Expiry.After(time.Now().Add(1*time.Minute)) {
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