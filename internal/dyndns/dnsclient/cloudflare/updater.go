package cloudflare

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type DNSUpdater struct {
	CloudflareConfig DNSRecordConfig
}

const apiPath = "https://api.cloudflare.com/client/v4/zones/"

func (receiver DNSUpdater) Update(ip string) (string, error) {
	fmt.Println(ip)
	config := receiver.CloudflareConfig
	payload, err := config.Payload(ip)

	if err != nil {
		return "", err
	}

	client := &http.Client{}
	url := fmt.Sprintf("%s%s/dns_records/%s", apiPath, config.ZoneIdentifier, config.DNSRecord.Identifier)

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(payload))
	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer " + config.Token},
	}
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result = string(body)

	return result, nil
}
