package ipproviders

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetIpProviderIpapi() IpProvider {
	return IpProvider{
		Name:          "Ipapi",
		ApiUrl:        "https://ipapi.co/json/",
		IsRateLimited: true,
		RateLimitType: MONTH_RATE_LIMIT,
		RateLimit:     1000,
		FuncGetIp:     getIpFromIpapi,
		IsEnabled:     true,
	}
}

func getIpFromIpapi() (string, error) {
	resp, err := http.Get("https://ipapi.co/json/")

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	return data["ip"].(string), nil
}
