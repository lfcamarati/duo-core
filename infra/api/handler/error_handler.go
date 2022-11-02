package handler

import (
	"github.com/gin-gonic/gin"
)

func ErrorHandler(fn func(ctx *gin.Context) ResponseError) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := fn(ctx)

		if err != nil {
			ctx.JSON(err.Code(), err)
		}
	}
}
