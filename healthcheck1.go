package main

import (
	"net/http"
	"time"

	"github.com/spf13/viper"
)

func HealthCheck1() HealthCheck {

	viper.SetDefault("healthcheck.hc1_url", "https://linuxctl.com/ip")
	url := viper.GetString("healthcheck.hc1_url")
	logger.Debug("HealthCheck1 called - using url:", url)

	completed := false
	status := 2

	hc := HealthCheck{
		Name:      "healthcheck1",
		Status:    status,
		Completed: completed,
		Date:      time.Now(),
		Verbose:   "testing",
	}

	resp, err := http.Get(url)
	if err != nil {
		hc.Status = 3
	} else {
		hc.Completed = true
		hc.Verbose = resp.Status
		if resp.StatusCode == 200 {
			hc.Status = 1
		}
	}
	defer resp.Body.Close()
	return hc

}
