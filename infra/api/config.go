package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()

	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Authorization,Content-Type,access-control-allow-origin,access-control-allow-headers"},
	}))

	InitRoutes(router)

	return router
}
