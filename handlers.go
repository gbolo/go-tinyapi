package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"math/rand"
	"os"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Index of tinyapi. These aren't the droids you're looking for\n")
}

func RouteHealthCheck(w http.ResponseWriter, r *http.Request) {
	hc1 := HealthCheck1()
	hc2 := HealthCheck2()
	hc := []HealthCheck{hc1, hc2}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if hc1.Status == 1 && hc2.Status == 1 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	if err := json.NewEncoder(w).Encode(hc); err != nil {
		panic(err)
	}
}

func RoutePanic(w http.ResponseWriter, r *http.Request) {
	logger.Fatal("RoutePanic invoked. Exiting")
	os.Exit(1)
}

func RouteChaosMonkey(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UTC().UnixNano())
	dice := rand.Intn(100)
	if dice > 90 {
		logger.Fatal("RouteChaosMonkey invoked. Exiting")
		os.Exit(1)
	} else {
		fmt.Fprintln(w, "Not Enough Chaos!", dice, "/100")
		logger.Warning("RouteChaosMonkey not enough chaos:", dice)
	}

}

func RouteOK(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "{ \"status\": \"OK\" }")
}