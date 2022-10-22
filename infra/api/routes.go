package api

import (
	"github.com/gin-gonic/gin"

	clientPfController "github.com/lfcamarati/duo-core/infra/api/controller/clientpf"
	clientPjController "github.com/lfcamarati/duo-core/infra/api/controller/clientpj"
	serviceController "github.com/lfcamarati/duo-core/infra/api/controller/service"
)

func InitRoutes(router *gin.Engine) {
	// ClientsPf
	router.GET("/clients-pf/:id", clientPfController.GetById)
	router.GET("/clients-pf", clientPfController.GetAll)
	router.POST("/clients-pf", clientPfController.Create)
	router.DELETE("/clients-pf/:id", clientPfController.Delete)

	// ClientsPj
	router.GET("/clients-pj/:id", clientPjController.GetById)
	router.GET("/clients-pj", clientPjController.GetAll)
	router.POST("/clients-pj", clientPjController.Create)
	router.DELETE("/clients-pj/:id", clientPjController.Delete)

	// Services
	router.GET("/services/:id", serviceController.GetById)
	router.GET("/services", serviceController.GetAll)
	router.POST("/services", serviceController.Create)
	router.DELETE("/services/:id", serviceController.Delete)
}
