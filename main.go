package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, "¡Bienvenido a la empresa Gophers!")
	})
	router.GET("/employees", func(c *gin.Context) {
		c.JSON(200, getEmployeeList())
	})
	router.GET("/employees/:id", func(c *gin.Context) {
		employeeId := c.Params.ByName("id")
		intVar, err := strconv.Atoi(employeeId)
		if err != nil {
			c.JSON(500, gin.H{"error": "Invalid id"})
		}
		e, er := getEmployeeById(intVar)
		if er != nil {
			c.JSON(404, gin.H{"error": er.Error()})
		} else {
			c.JSON(200, e)
		}
	})
	router.GET("/employeesparams", func(c *gin.Context) {
		valueId := c.Query("id")
		intVar, _ := strconv.Atoi(valueId)
		valueName := c.Query("name")
		valueActive := c.Query("active")
		boolVar, _ := strconv.ParseBool(valueActive)
		e := Employee{
			Id:     intVar,
			Name:   valueName,
			Active: boolVar,
		}
		newEmployeesList := addEmployee(e)
		c.JSON(201, newEmployeesList)
	})
	router.GET("/employeesactive", func(c *gin.Context) {
		c.JSON(200, getActiveEmployees(true))
	})

	router.Run(":8080")
}

type Employee struct {
	Id     int
	Name   string
	Active bool
}

func getActiveEmployees(isActive bool) any {
	var activeEmployees []Employee
	var noActiveEmployees []Employee
	employees := getEmployeeList()
	for _, e := range employees {
		if e.Active {
			activeEmployees = append(activeEmployees, e)
		} else {
			noActiveEmployees = append(noActiveEmployees, e)
		}
	}
	if isActive {
		return activeEmployees
	} else {
		return noActiveEmployees
	}
}

func addEmployee(e Employee) []Employee {
	employees := getEmployeeList()
	employees = append(employees, e)
	return employees
}

func getEmployeeList() []Employee {
	return []Employee{
		{Id: 1, Name: "John", Active: true},
		{Id: 2, Name: "Mary", Active: true},
		{Id: 3, Name: "Mike", Active: false},
		{Id: 4, Name: "Adam", Active: true},
		{Id: 5, Name: "Peter", Active: false},
	}
}

func getEmployeeById(id int) (Employee, error) {
	employees := getEmployeeList()
	for _, employee := range employees {
		if employee.Id == id {
			return employee, nil
		}
	}
	return Employee{}, fmt.Errorf("employee not found")
}
