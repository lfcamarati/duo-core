package main

import (
	"github.com/gin-gonic/gin"

	"github.com/lfcamarati/duo-core/cmd/routes"
	"github.com/lfcamarati/duo-core/pkg/database"
)

// https://github.com/golang-standards/project-layout/blob/master/README_ptBR.md

func main() {
	// Database
	database.DatabaseInit()

	// Gin
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// Routes
	routes.HandleRequest(router)

	// Start server
	router.Run()
}
