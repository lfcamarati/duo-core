package main

import (
	"github.com/lfcamarati/duo-core/infra/api"
	"github.com/lfcamarati/duo-core/infra/database"
)

// https://github.com/golang-standards/project-layout/blob/master/README_ptBR.md

func main() {
	// Database
	database.Init()

	// Http (Gin)
	router := api.Init()

	// Routes
	api.InitRoutes(router)

	// Start server
	router.Run()
}
