package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lfcamarati/duo-core/cmd/controller/client"
	"github.com/lfcamarati/duo-core/cmd/controller/service"
)

func HandleRequest(router *gin.Engine) {
	// Client
	router.POST("/clients", client.CreateClient)
	router.GET("/clients", client.GetAll)
	router.DELETE("/clients/:id", client.DeleteById)

	// Service
	router.GET("/services", service.GetAll)
}
