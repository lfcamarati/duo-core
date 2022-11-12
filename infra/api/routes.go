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
	var defaultHandlers = []handler.Handler{handler.SecurityHandler, handler.ErrorHandler}

	// Auth
	router.POST("/auth", handler.DefaultHandler(authController.Login, handler.ErrorHandler))
	router.GET("/auth", handler.DefaultHandler(authController.ValidateLogin, defaultHandlers...))

	// Users
	router.POST("/users", handler.DefaultHandler(userController.Create, defaultHandlers...))
	router.GET("/users/current", handler.DefaultHandler(userController.GetCurrent, defaultHandlers...))

	// Clients
	router.GET("/clients", handler.DefaultHandler(clientController.GetAll, defaultHandlers...))
	router.GET("/clients/:id", handler.DefaultHandler(clientController.GetById, defaultHandlers...))

	// ClientsPf
	router.GET("/clients-pf/:id", handler.DefaultHandler(clientPfController.GetById, defaultHandlers...))
	router.GET("/clients-pf", handler.DefaultHandler(clientPfController.GetAll, defaultHandlers...))
	router.POST("/clients-pf", handler.DefaultHandler(clientPfController.Create, defaultHandlers...))
	router.PUT("/clients-pf/:id", handler.DefaultHandler(clientPfController.Update, defaultHandlers...))
	router.DELETE("/clients-pf/:id", handler.DefaultHandler(clientPfController.Delete, defaultHandlers...))

	// ClientsPj
	router.GET("/clients-pj/:id", handler.DefaultHandler(clientPjController.GetById, defaultHandlers...))
	router.GET("/clients-pj", handler.DefaultHandler(clientPjController.GetAll, defaultHandlers...))
	router.POST("/clients-pj", handler.DefaultHandler(clientPjController.Create, defaultHandlers...))
	router.PUT("/clients-pj/:id", handler.DefaultHandler(clientPjController.Update, defaultHandlers...))
	router.DELETE("/clients-pj/:id", handler.DefaultHandler(clientPjController.Delete, defaultHandlers...))

	// Services
	router.GET("/services/:id", handler.DefaultHandler(serviceController.GetById, defaultHandlers...))
	router.GET("/services", handler.DefaultHandler(serviceController.GetAll, defaultHandlers...))
	router.POST("/services", handler.DefaultHandler(serviceController.Create, defaultHandlers...))
	router.PUT("/services/:id", handler.DefaultHandler(serviceController.Update, defaultHandlers...))
	router.DELETE("/services/:id", handler.DefaultHandler(serviceController.Delete, defaultHandlers...))
}
