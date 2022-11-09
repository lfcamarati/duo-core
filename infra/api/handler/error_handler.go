package handler

import (
	"github.com/gin-gonic/gin"
)

func ErrorHandler(fn ControllerEndpoint) ControllerEndpoint {
	return func(ctx *gin.Context) ResponseError {
		err := fn(ctx)

		if err != nil {
			ctx.JSON(err.Code(), err)
			return err
		}

		return nil
	}
}
