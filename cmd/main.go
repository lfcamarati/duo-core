package main

import (
	"github.com/lfcamarati/duo-core/infra/api"
	"github.com/lfcamarati/duo-core/infra/database"
	"github.com/lfcamarati/duo-core/infra/environment"
)

func main() {
	// Environment
	environment.Init()

	// Database
	database.Init()

	// Api (Gin)
	router := api.Init()

	// Start server
	router.Run()
}
