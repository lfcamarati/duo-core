package api

import (
	"github.com/gin-gonic/gin"

	clientController "github.com/lfcamarati/duo-core/infra/api/controller/client"
	clientPfController "github.com/lfcamarati/duo-core/infra/api/controller/clientpf"
	clientPjController "github.com/lfcamarati/duo-core/infra/api/controller/clientpj"
	serviceController "github.com/lfcamarati/duo-core/infra/api/controller/service"
	"github.com/lfcamarati/duo-core/infra/api/handler"
)

func InitRoutes(router *gin.Engine) {
	// Clients
	router.GET("/clients", clientController.GetAll)

	// ClientsPf
	router.GET("/clients-pf/:id", handler.ErrorHandler(clientPfController.GetById))
	router.GET("/clients-pf", handler.ErrorHandler(clientPfController.GetAll))
	router.POST("/clients-pf", handler.ErrorHandler(clientPfController.Create))
	router.PUT("/clients-pf/:id", handler.ErrorHandler(clientPfController.Update))
	router.DELETE("/clients-pf/:id", handler.ErrorHandler(clientPfController.Delete))

	// ClientsPj
	router.GET("/clients-pj/:id", clientPjController.GetById)
	router.GET("/clients-pj", clientPjController.GetAll)
	router.POST("/clients-pj", clientPjController.Create)
	router.PUT("/clients-pj/:id", clientPjController.Update)
	router.DELETE("/clients-pj/:id", clientPjController.Delete)

	// Services
	router.GET("/services/:id", serviceController.GetById)
	router.GET("/services", serviceController.GetAll)
	router.POST("/services", serviceController.Create)
	router.PUT("/services/:id", serviceController.Update)
	router.DELETE("/services/:id", serviceController.Delete)
}
