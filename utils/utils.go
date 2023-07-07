package utils

import (
	"fmt"
	"net/url"
	"regexp"
)

func IsMobileDevice(userAgent string) bool {
	if match, _ := regexp.MatchString("(?i)(Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini)", userAgent); match {
			return true
	}
	return false
}

func ParamsUrl(q string, baseURL string) string {
	params := url.Values{}
	params.Set("q", q)
	params.Set("type", "track")
	params.Set("market", "JP")
	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	
	return fullURL
}
