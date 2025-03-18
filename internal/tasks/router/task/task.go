package taskrouter

import "github.com/gin-gonic/gin"

func Register(router *gin.Engine) {
	v1 := router.Group("/v1/task")
	{
		v1.POST("")
	}
}
