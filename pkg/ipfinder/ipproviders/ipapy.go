package ipproviders

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetIpProviderIpapi() IpProvider {
	return IpProvider{
		Name:          "Ipapi",
		ApiUrl:        "https://ipapi.co/json/",
		IsRateLimited: false,
		RateLimitType: SECOND_RATE_LIMIT,
		RateLimit:     0,
		FuncGetIp:     getIpFromIpapi,
		IsEnabled:     true,
	}
}

func getIpFromIpapi() string {
	resp, err := http.Get("https://ipapi.co/json/")

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println(err)
	}

	return data["ip"].(string)
}
