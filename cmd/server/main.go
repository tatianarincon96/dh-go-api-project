package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	handlers2 "github.com/tatianarincon96/dh-go-api-project/cmd/server/handlers"
	"github.com/tatianarincon96/dh-go-api-project/cmd/server/internal/domain"
	"github.com/tatianarincon96/dh-go-api-project/cmd/server/store"
	"log"
	"os"
)

var productList []domain.Product
var employeeList []domain.Employee
var lastProductID int

func main() {
	// Get envvar
	e := godotenv.Load()
	if e != nil {
		log.Fatal("Error loading .env file")
	}
	usuario := os.Getenv("USER")
	contrasena := os.Getenv("PASSWORD")
	fmt.Printf("Usuario: %s, Contraseña: %s", usuario, contrasena)

	// Crear router
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// Definir router de chequeo
	handlers2.NewHealthRouter(router)

	// Inicializar productos
	productList = store.LoadProducts("../../products.json")
	employeeList = store.LoadEmployees("../../employees.json")
	lastProductID = len(productList)
	handlers2.NewProductRouter(router, productList, lastProductID)
	handlers2.NewEmployeeRouter(router, employeeList)

	// Iniciar el servidor
	router.Run(":8080")
}
