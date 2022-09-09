package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	handlers2 "github.com/tatianarincon96/dh-go-api-project/handlers"
	"log"
	"os"
)

// @title PRODUCTS Service API
// @version 1.0.0
// @description This API handles Supermarket products.
// @termsOfService http://localhost:8080/terms

// contact.name Tatiana Rincon
// contact.email

// lince.name Apache 2.0
func main() {
	// Crear router
	router := gin.Default()
	/*router.SetTrustedProxies([]string{"127.0.0.1"})*/

	// ADD MIDDLEWARE
	router.Use(TokenAuthMiddleware())

	// Definir router de chequeo
	handlers2.NewHealthRouter(router)
	handlers2.NewProductRouter(router, *handlers2.InitProductHandler())
	/*handler.NewEmployeeRouter(router, employeeList)*/

	// Iniciar el servidor
	router.Run(":8080")
}

func TokenAuthMiddleware() gin.HandlerFunc {
	e := godotenv.Load()
	if e != nil {
		log.Fatal("Error loading .env file")
	}
	requiredToken := os.Getenv("API_TOKEN")
	return func(c *gin.Context) {
		token := c.GetHeader("api_token")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "API token required"})
			return
		}
		if token != requiredToken {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid API token"})
			return
		}
		c.Next()
	}
}
