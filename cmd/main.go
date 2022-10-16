package main

import (
	"github.com/gin-contrib/cors"
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
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))

	// Routes
	routes.HandleRequest(router)

	// Start server
	router.Run()
}
