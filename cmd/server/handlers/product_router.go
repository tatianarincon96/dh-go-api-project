package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tatianarincon96/dh-go-api-project/cmd/server/internal/domain"
	"github.com/tatianarincon96/dh-go-api-project/cmd/server/internal/product"
	"strconv"
)

func NewProductRouter(e *gin.Engine, list []domain.Product) {
	products := e.Group("/api/v1/product")
	{
		products.GET("", GetAllProducts(list))
		products.GET("/:id", GetProductById(list))
		products.GET("/searchbyquantity", SearchProduct(list))
		products.GET("/buy", BuyProduct(list))
		products.GET("/add", AddProduct(list))
	}
}

func GetAllProducts(list []domain.Product) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, list)
	}
}

func GetProductById(list []domain.Product) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid id"})
		}
		p, e := product.GetProductById(list, id)
		if e != nil {
			c.JSON(404, gin.H{"error": e})
		}
		c.JSON(200, p)
	}
}

func SearchProduct(list []domain.Product) gin.HandlerFunc {
	return func(c *gin.Context) {
		minLimit := c.Query("min_limit")
		intMinLimit, _ := strconv.Atoi(minLimit)
		maxLimit := c.Query("max_limit")
		intMaxLimit, _ := strconv.Atoi(maxLimit)
		c.JSON(200, product.GetProductsByQuantity(list, intMinLimit, intMaxLimit))
	}
}

func BuyProduct(list []domain.Product) gin.HandlerFunc {
	return func(c *gin.Context) {
		codeValue := c.Query("code_value")
		quantityToBuy := c.Query("quantity_to_buy")
		intQuantityToBuy, _ := strconv.Atoi(quantityToBuy)
		c.JSON(200, product.BuyProduct(list, codeValue, intQuantityToBuy))
	}
}

func AddProduct(list []domain.Product) gin.HandlerFunc {
	return func(c *gin.Context) {
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
		p := domain.Product{
			Id:          intId,
			Name:        valueName,
			Quantity:    intQuantity,
			CodeValue:   valueCodeValue,
			IsPublished: boolVar,
			Expiration:  valueExpiration,
			Price:       floatVar,
		}
		product.AddProduct(list, p)
		c.JSON(200, gin.H{"product": p})
	}
}
