package main

import (
	"net/http"
)

func httpServer() error {

	loadConfig()

	logger.Info("Starting server on:", getServerConfig())

	router := NewRouter()
	return http.ListenAndServe(getServerConfig(), router)

}
