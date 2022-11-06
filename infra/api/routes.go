package api

import (
	"github.com/gin-gonic/gin"

	authController "github.com/lfcamarati/duo-core/infra/api/controller/auth"
	clientController "github.com/lfcamarati/duo-core/infra/api/controller/client"
	clientPfController "github.com/lfcamarati/duo-core/infra/api/controller/clientpf"
	clientPjController "github.com/lfcamarati/duo-core/infra/api/controller/clientpj"
	serviceController "github.com/lfcamarati/duo-core/infra/api/controller/service"
	userController "github.com/lfcamarati/duo-core/infra/api/controller/user"
	"github.com/lfcamarati/duo-core/infra/api/handler"
)

func InitRoutes(router *gin.Engine) {
	// Auth
	router.POST("/auth", handler.ErrorHandler(authController.Login))
	router.GET("/auth", handler.DefaultHandler(authController.ValidateLogin))

	// Users
	router.POST("/users", handler.DefaultHandler(userController.Create))

	// Clients
	router.GET("/clients", handler.DefaultHandler(clientController.GetAll))

	// ClientsPf
	router.GET("/clients-pf/:id", handler.DefaultHandler(clientPfController.GetById))
	router.GET("/clients-pf", handler.DefaultHandler(clientPfController.GetAll))
	router.POST("/clients-pf", handler.DefaultHandler(clientPfController.Create))
	router.PUT("/clients-pf/:id", handler.DefaultHandler(clientPfController.Update))
	router.DELETE("/clients-pf/:id", handler.DefaultHandler(clientPfController.Delete))

	// ClientsPj
	router.GET("/clients-pj/:id", handler.DefaultHandler(clientPjController.GetById))
	router.GET("/clients-pj", handler.DefaultHandler(clientPjController.GetAll))
	router.POST("/clients-pj", handler.DefaultHandler(clientPjController.Create))
	router.PUT("/clients-pj/:id", handler.DefaultHandler(clientPjController.Update))
	router.DELETE("/clients-pj/:id", handler.DefaultHandler(clientPjController.Delete))

	// Services
	router.GET("/services/:id", handler.DefaultHandler(serviceController.GetById))
	router.GET("/services", handler.DefaultHandler(serviceController.GetAll))
	router.POST("/services", handler.DefaultHandler(serviceController.Create))
	router.PUT("/services/:id", handler.DefaultHandler(serviceController.Update))
	router.DELETE("/services/:id", handler.DefaultHandler(serviceController.Delete))
}
