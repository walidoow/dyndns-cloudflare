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
	SECOND_RATE_LIMIT = 1
	MINUTE_RATE_LIMIT = 60 * SECOND_RATE_LIMIT
	HOUR_RATE_LIMIT   = 60 * MINUTE_RATE_LIMIT
	DAY_RATE_LIMIT    = 24 * HOUR_RATE_LIMIT
	WEEK_RATE_LIMIT   = 7 * DAY_RATE_LIMIT
	MONTH_RATE_LIMIT  = 30 * DAY_RATE_LIMIT
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
	PeriodSizeSec float32
	FuncGetIp     GetIp
	IsEnabled     bool
	LastUsage     int64
}
