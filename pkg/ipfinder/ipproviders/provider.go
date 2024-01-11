package ipproviders

var IpProviders = []IpProvider{
	GetIpProviderIpify(),
	GetIpProviderIpapi(),
	GetIpProviderIcanhazip(),
	GetIpProviderBigdatacloud(),
}

type RateLimitType int

type GetIp func() string

const (
	SECOND_RATE_LIMIT = 0
	MINUTE_RATE_LIMIT = 1
	HOUR_RATE_LIMIT   = 2
	DAY_RATE_LIMIT    = 3
	WEEK_RATE_LIMIT   = 4
	MONTH_RATE_LIMIT  = 5
)

type IpProvider struct {
	Name          string
	ApiUrl        string
	IsRateLimited bool
	RateLimitType RateLimitType
	RateLimit     uint
	FuncGetIp     GetIp
	IsEnabled     bool
}
