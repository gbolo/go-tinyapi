package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"RouteHealthCheck",
		"GET",
		"/healthcheck",
		RouteHealthCheck,
	},
	Route{
		"RoutePanic",
		"GET",
		"/panic",
		RoutePanic,
	},
	Route{
		"RouteChaosMonkey",
		"GET",
		"/chaosmonkey",
		RouteChaosMonkey,
	},
	Route{
		"AlwaysOK",
		"GET",
		"/alwaysok",
		RouteOK,
	},
}
