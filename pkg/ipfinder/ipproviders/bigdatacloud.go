package ipproviders

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetIpProviderBigdatacloud() IpProvider {
	return IpProvider{
		Name:          "Bigdatacloud",
		ApiUrl:        "https://api-bdc.net/data/client-ip",
		IsRateLimited: false,
		RateLimitType: SECOND_RATE_LIMIT,
		RateLimit:     0,
		FuncGetIp:     getIpFromBigdatacloud,
		IsEnabled:     true,
	}
}

func getIpFromBigdatacloud() (string, error) {
	resp, err := http.Get("https://api-bdc.net/data/client-ip")

	if err != nil {
		fmt.Println(err)
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

	return data["ipString"].(string), nil
}
