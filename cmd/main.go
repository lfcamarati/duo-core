package main

import (
	"github.com/lfcamarati/duo-core/infra/api"
	"github.com/lfcamarati/duo-core/infra/database"
)

func main() {
	// Database
	database.Init()

	// Http (Gin)
	router := api.Init()

	// Start server
	router.Run()
}
