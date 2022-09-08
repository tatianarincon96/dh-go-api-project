package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type errorResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Usar generics
type response struct {
	Data interface{} `json:"data"`
}

func Success(ctx *gin.Context, status int, data interface{}) {
	ctx.JSON(status, response{Data: data})
}

func Failure(ctx *gin.Context, status int, err error) {
	ctx.JSON(status, errorResponse{
		Status:  status,
		Code:    http.StatusText(status),
		Message: err.Error(),
	})
}
