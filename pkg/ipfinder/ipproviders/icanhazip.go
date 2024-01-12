package ipproviders

import (
	"io"
	"net/http"
)

func GetIpProviderIcanhazip() IpProvider {
	return IpProvider{
		Name:          "Icanhazip",
		ApiUrl:        "https://icanhazip.com/",
		IsRateLimited: false,
		RateLimitType: SECOND_RATE_LIMIT,
		RateLimit:     0,
		FuncGetIp:     getIpFromIcanhazip,
		IsEnabled:     true,
	}
}

func getIpFromIcanhazip() (string, error) {
	resp, err := http.Get("https://icanhazip.com/")

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body[0:15]), nil // Exclude last byte which is a new line character
}
