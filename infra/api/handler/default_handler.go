package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lfcamarati/duo-core/infra/security"
)

func DefaultHandler(fn func(ctx *gin.Context) ResponseError) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.Request.Header["Authorization"]

		if len(authorizationHeader) == 0 {
			ctx.Status(http.StatusUnauthorized)
			return
		}

		authErr := security.VerifyJWT(authorizationHeader[0])

		if authErr != nil {
			ctx.Status(http.StatusUnauthorized)
			return
		}

		err := fn(ctx)

		if err != nil {
			ctx.JSON(err.Code(), err)
		}
	}
}
