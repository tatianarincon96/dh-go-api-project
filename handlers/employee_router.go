package handlers

import (
	"github.com/gin-gonic/gin"
)

func NewEmployeeRouter(e *gin.Engine, h employeeHandler) {
	// Validar token de forma central con middleware
	employees := e.Group("/api/v1/employees")
	{
		employees.GET("", h.GetAll())
		employees.GET("/:id", h.GetById())
		employees.POST("", h.Create())
		employees.PUT("", h.Update())
		employees.DELETE("/:id", h.DeleteByID())
		employees.PATCH("", h.UpdateField())
	}
}
