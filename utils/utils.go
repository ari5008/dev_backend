package utils

import (
	"fmt"
	"net/url"
)

func ParamsUrl(q string, baseURL string) string {
	params := url.Values{}
	params.Set("q", q)
	params.Set("type", "track")
	params.Set("market", "JP")
	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	
	return fullURL
}
