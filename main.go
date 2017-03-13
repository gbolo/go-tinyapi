package main

// Constants go here.
const configName = "config"
const appName = "tinyapi"
const version = "v0.1"

// Start main()
func main() {

	initLogging()

	logger.Info("[STARTUP] Initializing", appName, version)
	logger.Fatal(httpServer())
}
