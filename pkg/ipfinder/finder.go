package ipfinder

import (
	"dyndns-cloudflare/pkg/ipfinder/ipproviders"
	"time"
)

// ApiClient is a struct that holds the IP Providers and the index of the next IP Provider to use
type ApiClient struct {
	ipProviders []ipproviders.IpProvider
	index       int
}

// New returns a new ApiClient
func New() *ApiClient {
	client := ApiClient{index: 0}

	for _, ipProvider := range ipproviders.IpProviders {
		if ipProvider.IsEnabled {
			client.ipProviders = append(client.ipProviders, ipProvider)
		}
	}
	return &client
}

// GetCurrentIp returns the current IP from the next available IP Provider
func (ac *ApiClient) GetCurrentIp() string {
	var ipProvider ipproviders.IpProvider = ac.ipProviders[ac.index]

	response, err := ipProvider.FuncGetIp()
	if err != nil {
		return "" // TODO: handle error
	}
	ipProvider.LastUsage = time.Now().Unix()

	ac.index = (ac.index + 1) % len(ac.ipProviders)
	return response
}
