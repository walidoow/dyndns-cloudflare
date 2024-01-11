package ipproviders

import (
	"io"
	"log"
	"net/http"
)

func GetIpProviderIpify() IpProvider {
	return IpProvider{
		Name:          "Ipify",
		ApiUrl:        "https://api.ipify.org",
		IsRateLimited: false,
		RateLimitType: SECOND_RATE_LIMIT,
		RateLimit:     0,
		FuncGetIp:     getIpFromIpify,
		IsEnabled:     true,
	}
}

func getIpFromIpify() string {
	resp, err := http.Get("https://api.ipify.org")

	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	return string(body)
}
