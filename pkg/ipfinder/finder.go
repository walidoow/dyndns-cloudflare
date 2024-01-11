package ipfinder

import "dyndns-cloudflare/pkg/ipfinder/ipproviders"

type ApiClient struct {
	ipProviders []ipproviders.IpProvider
	index       int
}

func New() *ApiClient {
	client := ApiClient{index: 0}

	for _, ipProvider := range ipproviders.IpProviders {
		if ipProvider.IsEnabled {
			client.ipProviders = append(client.ipProviders, ipProvider)
		}
	}
	return &client
}

func (ac *ApiClient) GetCurrentIp() string {
	response := ac.ipProviders[ac.index].FuncGetIp()
	ac.index = (ac.index + 1) % len(ac.ipProviders)
	return response
}
