package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lfcamarati/duo-core/cmd/controller/client"
	"github.com/lfcamarati/duo-core/cmd/controller/service"
)

func HandleRequest(router *gin.Engine) {
	// Client
	router.POST("/clients", client.NewClientPf)
	router.GET("/clients", client.GetAll)

	// Service
	router.GET("/services", service.GetAll)
}
