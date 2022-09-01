package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tatianarincon96/dh-go-api-project/cmd/server/internal/domain"
	"github.com/tatianarincon96/dh-go-api-project/cmd/server/internal/employee"
	"strconv"
)

func NewEmployeeRouter(e *gin.Engine, list []domain.Employee) {
	employee := e.Group("/api/v1/employee")
	{
		employee.GET("", GetAllEmployees(list))
		employee.GET("/:id", GetEmployeeById(list))
		employee.GET("/active", ActiveEmployees(list))
		employee.GET("/add", AddEmployee(list))
	}
}

func GetAllEmployees(list []domain.Employee) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, list)
	}
}

func GetEmployeeById(list []domain.Employee) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid id"})
		}
		p, e := employee.GetEmployeeById(list, id)
		if e != nil {
			c.JSON(404, gin.H{"error": e})
		}
		c.JSON(200, p)
	}
}

func ActiveEmployees(list []domain.Employee) gin.HandlerFunc {
	return func(c *gin.Context) {
		checkActive := c.Query("check_active")
		boolCheckActive, _ := strconv.ParseBool(checkActive)
		c.JSON(200, employee.GetActiveEmployees(list, boolCheckActive))
	}
}

func AddEmployee(list []domain.Employee) gin.HandlerFunc {
	return func(c *gin.Context) {
		valueId := c.Query("id")
		intVar, _ := strconv.Atoi(valueId)
		valueName := c.Query("name")
		valueActive := c.Query("active")
		boolVar, _ := strconv.ParseBool(valueActive)
		e := domain.Employee{
			Id:     intVar,
			Name:   valueName,
			Active: boolVar,
		}
		newEmployeesList := employee.AddEmployee(list, e)
		c.JSON(200, newEmployeesList)
	}
}
