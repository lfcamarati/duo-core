package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lfcamarati/duo-core/infra/security"
)

func SecurityHandler(fn ControllerEndpoint) ControllerEndpoint {
	return func(ctx *gin.Context) ResponseError {
		authorizationHeader := ctx.Request.Header["Authorization"]

		if len(authorizationHeader) == 0 {
			ctx.Status(http.StatusUnauthorized)
			return nil
		}

		authErr := security.VerifyJWT(authorizationHeader[0])

		if authErr != nil {
			ctx.Status(http.StatusUnauthorized)
			return nil
		}

		return fn(ctx)
	}
}
