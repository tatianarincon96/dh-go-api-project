package main

import (
	"github.com/gin-gonic/gin"
	handlers2 "github.com/tatianarincon96/dh-go-api-project/cmd/server/handlers"
	"github.com/tatianarincon96/dh-go-api-project/cmd/server/store"
)

func main() {
	// Crear router
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// Definir router de chequeo
	handlers2.NewHealthRouter(router)

	// Inicializar productos
	productList := store.LoadProducts("./product.json")
	employeeList := store.LoadEmployees("./employee.json")
	handlers2.NewProductRouter(router, productList)
	handlers2.NewEmployeeRouter(router, employeeList)

	// Iniciar el servidor
	router.Run(":8080")
}
