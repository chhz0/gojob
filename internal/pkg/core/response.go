package core

import (
	"net/http"

	"github.com/chhz0/gojob/internal/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Caused  string `json:"cause,omitempty"`
	Message string `json:"message,omitempty"`
}

func WriteResponse(c *gin.Context, err error, data interface{}) {
	if err != nil {
		err := errcode.From(err)
		c.JSON(err.Code, gin.H{
			"caused":  err.Caused,
			"message": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, data)
}
