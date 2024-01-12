package cloudflare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type DNSUpdater struct {
	CloudflareConfig DNSRecordConfig
}

const apiPath = "https://api.cloudflare.com/client/v4/zones/"

func (receiver DNSUpdater) Update(ip string) (string, error) {
	config := receiver.CloudflareConfig
	payload, err := config.Payload(ip)

	if err != nil {
		return "", err
	}

	client := &http.Client{}
	url := fmt.Sprintf("%s%s/dns_records/%s", apiPath, config.ZoneIdentifier, config.DNSRecord.Identifier)

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}

	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer " + config.Token},
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result = string(body)

	return result, nil
}

func (receiver DNSUpdater) DNSRecordIp() (string, error) {
	config := receiver.CloudflareConfig

	client := &http.Client{}
	url := fmt.Sprintf("%s%s/dns_records/%s", apiPath, config.ZoneIdentifier, config.DNSRecord.Identifier)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", nil
	}
	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer " + config.Token},
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response ApiGetRecordResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	var result = response.Result.Content

	return result, nil
}
