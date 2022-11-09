package handler

import (
	"github.com/gin-gonic/gin"
)

type ControllerEndpoint func(ctx *gin.Context) ResponseError
type Handler func(ControllerEndpoint) ControllerEndpoint

func DefaultHandler(fn ControllerEndpoint, handlers ...Handler) gin.HandlerFunc {
	if len(handlers) < 1 {
		return wrap(fn)
	}

	wrapped := fn

	for i := len(handlers) - 1; i >= 0; i-- {
		wrapped = handlers[i](wrapped)
	}

	return wrap(wrapped)
}

func wrap(fn ControllerEndpoint) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fn(ctx)
	}
}
