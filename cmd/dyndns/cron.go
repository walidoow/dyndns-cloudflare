package main

import (
	"dyndns-cloudflare/pkg/ipfinder"
	"fmt"
	"github.com/robfig/cron"
	"log"
)

var cachedDNSProviderIp string

func run(clientApi *ipfinder.ApiClient) {
	currentIp := clientApi.GetCurrentIp()

	if cachedDNSProviderIp != currentIp {
		log.Println("Updating the cloudflare record ip from " + cachedDNSProviderIp + " to " + currentIp)
		_, err := updater.Update(currentIp)
		if err != nil {
			log.Println("Error while trying yo update ip on the DNS Record")
			return
		}
		cachedDNSProviderIp = currentIp
	}
}

func RunCronJobs(clientApi *ipfinder.ApiClient) {
	c := cron.New()

	var err error
	cachedDNSProviderIp, err = updater.DNSRecordIp()
	if err != nil {
		log.Fatal(err)
	}

	err = c.AddFunc("@every 1s", func() { run(clientApi) })
	if err != nil {
		log.Fatal(err)
	}

	c.Start()

	_, err = fmt.Scanln()
	if err != nil {
		log.Fatal(err)
	}
}
