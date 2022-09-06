package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	handlers2 "github.com/tatianarincon96/dh-go-api-project/handlers"
	"log"
	"os"
)

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
	/*router.SetTrustedProxies([]string{"127.0.0.1"})*/

	// Definir router de chequeo
	handlers2.NewHealthRouter(router)
	handlers2.NewProductRouter(router, *handlers2.InitProductHandler())
	/*handler.NewEmployeeRouter(router, employeeList)*/

	// Iniciar el servidor
	router.Run(":8080")
}
