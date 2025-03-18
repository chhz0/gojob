package router

import (
	"net/http"

	taskrouter "github.com/chhz0/gojob/internal/job/router/task"
	"github.com/chhz0/gojob/internal/pkg/core"
	"github.com/chhz0/gojob/internal/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Page not found.",
		})
		core.WriteResponse(ctx, errcode.ErrNotFound.WithMessage("Page not found."), nil)
	})

	router.GET("/healthz", func(ctx *gin.Context) {
		core.WriteResponse(ctx, nil, gin.H{
			"status": "ok",
		})
	})

	taskrouter.Register(router)
}
