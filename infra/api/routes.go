package api

import (
	"github.com/gin-gonic/gin"

	clientController "github.com/lfcamarati/duo-core/infra/api/controller/client"
	serviceController "github.com/lfcamarati/duo-core/infra/api/controller/service"
)

func InitRoutes(router *gin.Engine) {
	// Client
	router.POST("/clients", clientController.CreateClient)
	router.GET("/clients", clientController.GetAll)
	router.DELETE("/clients/:id", clientController.DeleteById)

	// Service
	router.GET("/services", serviceController.GetAll)
}
