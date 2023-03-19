package api

import (
	"github.com/gin-gonic/gin"

	authController "github.com/lfcamarati/duo-core/infra/api/controller/auth"
	clientController "github.com/lfcamarati/duo-core/infra/api/controller/client"
	serviceController "github.com/lfcamarati/duo-core/infra/api/controller/service"
	userController "github.com/lfcamarati/duo-core/infra/api/controller/user"
	"github.com/lfcamarati/duo-core/infra/api/handler"
)

func InitRoutes(router *gin.Engine) {
	var defaultHandlers = []handler.Handler{handler.SecurityHandler, handler.ErrorHandler}

	routerGroup := router.Group("/api")

	// Auth
	routerGroup.POST("/auth", handler.DefaultHandler(authController.Login, handler.ErrorHandler))
	routerGroup.GET("/auth", handler.DefaultHandler(authController.ValidateLogin, defaultHandlers...))

	// Users
	routerGroup.POST("/users", handler.DefaultHandler(userController.Create, defaultHandlers...))
	routerGroup.GET("/users/current", handler.DefaultHandler(userController.GetCurrent, defaultHandlers...))

	// Clients
	routerGroup.GET("/clients", handler.DefaultHandler(clientController.GetAll, defaultHandlers...))
	routerGroup.GET("/clients/:id", handler.DefaultHandler(clientController.GetById, defaultHandlers...))
	routerGroup.POST("/clients", handler.DefaultHandler(clientController.Create, defaultHandlers...))
	routerGroup.PUT("/clients/:id", handler.DefaultHandler(clientController.Update, defaultHandlers...))
	routerGroup.DELETE("/clients/:id", handler.DefaultHandler(clientController.Delete, defaultHandlers...))

	// Services
	routerGroup.GET("/services/:id", handler.DefaultHandler(serviceController.GetById, defaultHandlers...))
	routerGroup.GET("/services", handler.DefaultHandler(serviceController.GetAll, defaultHandlers...))
	routerGroup.POST("/services", handler.DefaultHandler(serviceController.Create, defaultHandlers...))
	routerGroup.PUT("/services/:id", handler.DefaultHandler(serviceController.Update, defaultHandlers...))
	routerGroup.DELETE("/services/:id", handler.DefaultHandler(serviceController.Delete, defaultHandlers...))
}
