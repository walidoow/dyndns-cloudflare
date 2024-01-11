package main

import (
	"dyndns-cloudflare/internal/dyndns/dnsclient"
	"dyndns-cloudflare/internal/dyndns/dnsclient/cloudflare"
	"dyndns-cloudflare/pkg/ipfinder"
)

var updater dnsclient.DNSUpdater

func main() {
	var clientApi = ipfinder.New()
	var provider = CLOUDFLARE
	var configPath = "config.yml"

	switch provider {
	case CLOUDFLARE:
		updater = cloudflare.DNSUpdater{CloudflareConfig: cloudflare.LoadYAMLConfig(configPath)}
	}

	RunCronJobs(clientApi)
}
