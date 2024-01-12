package ipproviders

// IpProviders is a slice of all the available IP providers
var IpProviders = []IpProvider{
	GetIpProviderIpify(),
	GetIpProviderIpapi(),
	GetIpProviderIcanhazip(),
	GetIpProviderBigdatacloud(),
}

// RateLimitType is an enum that represents the type of rate limit
type RateLimitType int

const (
	SECOND_RATE_LIMIT = 0
	MINUTE_RATE_LIMIT = 1
	HOUR_RATE_LIMIT   = 2
	DAY_RATE_LIMIT    = 3
	WEEK_RATE_LIMIT   = 4
	MONTH_RATE_LIMIT  = 5
)

// GetIp is a function that returns the current IP
type GetIp func() (string, error)

// IpProvider is a struct that contains all the information needed to get the current IP from a provider
type IpProvider struct {
	Name          string
	ApiUrl        string
	IsRateLimited bool
	RateLimitType RateLimitType
	RateLimit     uint
	FuncGetIp     GetIp
	IsEnabled     bool
	LastUsage     int64
}
