package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Teste struct {
	Nome string
}

func GetAll(c *gin.Context) {
	teste := Teste{Nome: "Service A"}
	c.JSON(http.StatusOK, teste)
}
