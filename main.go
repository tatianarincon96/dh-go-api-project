package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tatianarincon96/dh-go-api-project/employees"
	"github.com/tatianarincon96/dh-go-api-project/products"
	"strconv"
)

func main() {
	// Crear router
	router := gin.Default()

	// Home route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, "¡Bienvenido a la empresa Gophers!")
	})

	// Endpoint de Empleados
	router.GET("/employees", func(c *gin.Context) {
		c.JSON(200, employees.GetEmployeeList())
	})
	router.GET("/employees/:id", func(c *gin.Context) {
		employeeId := c.Params.ByName("id")
		intVar, err := strconv.Atoi(employeeId)
		if err != nil {
			c.JSON(500, gin.H{"error": "Invalid id"})
		}
		e, er := employees.GetEmployeeById(intVar)
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
		e := employees.Employee{
			Id:     intVar,
			Name:   valueName,
			Active: boolVar,
		}
		newEmployeesList := employees.AddEmployee(e)
		c.JSON(200, newEmployeesList)
	})
	router.GET("/employeesactive", func(c *gin.Context) {
		c.JSON(200, employees.GetActiveEmployees(true))
	})

	// Endpoint de Productos
	router.GET("/productparams", func(c *gin.Context) {
		valueId := c.Query("id")
		intId, _ := strconv.Atoi(valueId)
		valueName := c.Query("name")
		valueQuantity := c.Query("quantity")
		intQuantity, _ := strconv.Atoi(valueQuantity)
		valueCodeValue := c.Query("code_value")
		valueIsPublished := c.Query("is_published")
		boolVar, _ := strconv.ParseBool(valueIsPublished)
		valueExpiration := c.Query("expiration")
		valuePrice := c.Query("price")
		floatVar, _ := strconv.ParseFloat(valuePrice, 64)
		p := products.Product{
			Id:          intId,
			Name:        valueName,
			Quantity:    intQuantity,
			CodeValue:   valueCodeValue,
			IsPublished: boolVar,
			Expiration:  valueExpiration,
			Price:       floatVar,
		}
		products.AddProduct(p)
		c.JSON(200, gin.H{"product": p})
	})
	router.GET("/products/:id", func(c *gin.Context) {
		productId := c.Params.ByName("id")
		intVar, err := strconv.Atoi(productId)
		if err != nil {
			c.JSON(500, gin.H{"error": "Invalid id"})
		}
		e, er := products.GetProductById(intVar)
		if er != nil {
			c.JSON(404, gin.H{"error": er.Error()})
		} else {
			c.JSON(200, e)
		}
	})
	router.GET("/products/searchbyquantity", func(c *gin.Context) {
		minLimit := c.Query("min_limit")
		intMinLimit, _ := strconv.Atoi(minLimit)
		maxLimit := c.Query("max_limit")
		intMaxLimit, _ := strconv.Atoi(maxLimit)
		c.JSON(200, products.GetProductsByQuantity(intMinLimit, intMaxLimit))
	})
	router.GET("/products/buy", func(c *gin.Context) {
		codeValue := c.Query("code_value")
		quantityToBuy := c.Query("quantity_to_buy")
		intQuantityToBuy, _ := strconv.Atoi(quantityToBuy)
		c.JSON(200, products.BuyProduct(codeValue, intQuantityToBuy))
	})

	// Iniciar el servidor
	router.Run(":8080")
}
