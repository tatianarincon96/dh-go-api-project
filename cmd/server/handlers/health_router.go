package handlers

import "github.com/gin-gonic/gin"

func NewHealthRouter(e *gin.Engine) {
	e.GET("/", func(c *gin.Context) {
		c.JSON(200, "¡Bienvenido a la empresa Gophers!")
	})
	e.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})
}
