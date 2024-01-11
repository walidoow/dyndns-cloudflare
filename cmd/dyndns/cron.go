package main

import (
	"dyndns-cloudflare/pkg/ipfinder"
	"fmt"
	"github.com/robfig/cron"
	"log"
)

func run(clientApi *ipfinder.ApiClient) {
	currentIp := clientApi.GetCurrentIp()
	updater.Update(currentIp)
}

func RunCronJobs(clientApi *ipfinder.ApiClient) {
	c := cron.New()

	err := c.AddFunc("@every 1s", func() { run(clientApi) })
	if err != nil {
		log.Fatal(err)
	}

	c.Start()

	_, err = fmt.Scanln()
	if err != nil {
		log.Fatal(err)
	}
}
