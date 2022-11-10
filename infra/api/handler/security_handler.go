package handler

import (
	"net/http"
	"strings"

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

		tokenString := strings.Replace(authorizationHeader[0], "Bearer ", "", 1)
		authErr := security.VerifyJWT(tokenString)

		if authErr != nil {
			ctx.Status(http.StatusUnauthorized)
			return nil
		}

		return fn(ctx)
	}
}
