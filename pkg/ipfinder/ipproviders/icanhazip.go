package ipproviders

import (
	"io"
	"log"
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

func getIpFromIcanhazip() string {
	resp, err := http.Get("https://icanhazip.com/")

	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
	}

	return string(body[0:15]) // Exclude last byte which is a new line character
}
