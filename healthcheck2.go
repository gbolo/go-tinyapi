package main

import (
	"net/http"
	"time"

	"github.com/spf13/viper"
)

func HealthCheck2() HealthCheck {

	viper.SetDefault("healthcheck.hc2_url", "https://google.com/some404")
	url := viper.GetString("healthcheck.hc2_url")
	logger.Debug("HealthCheck2 called - using url:", url)

	completed := false
	status := 2

	hc := HealthCheck{
		Name:      "healthcheck2",
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
