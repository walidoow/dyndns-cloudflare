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
			// Calculated field for rate limiting
			ipProvider.PeriodSizeSec = float32(ipProvider.RateLimit) / float32(ipProvider.RateLimitType)
			client.ipProviders = append(client.ipProviders, ipProvider)
		}
	}
	return &client
}

// GetCurrentIp returns the current IP from the next available IP Provider
func (ac *ApiClient) GetCurrentIp() string {
	currentIpProvider, err := selectNextAvailableIpProvider(ac)
	if err != nil {
		return "" // TODO: handle error
	}

	response, err := currentIpProvider.FuncGetIp()
	if err != nil {
		return "" // TODO: handle error
	}
	currentIpProvider.LastUsage = time.Now().Unix()

	return response
}

// selectNextAvailableIpProvider returns the next available IP Provider.
// If all IP Providers are rate limited, it returns an error.
func selectNextAvailableIpProvider(ac *ApiClient) (ipproviders.IpProvider, error) {
	for i := 0; i < len(ac.ipProviders); i++ {
		ac.index = ac.calculateNextIndex()
		ipProvider := ac.ipProviders[ac.index]
		if !ipProvider.IsRateLimited || time.Now().Unix()-ipProvider.LastUsage >= int64(ipProvider.PeriodSizeSec) {
			return ipProvider, nil
		}
	}
	return ipproviders.IpProvider{}, nil
}

// calculateNextIndex returns the next index in the slice of IP Providers
func (ac *ApiClient) calculateNextIndex() int {
	return (ac.index + 1) % len(ac.ipProviders)
}
