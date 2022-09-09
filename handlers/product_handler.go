package handlers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tatianarincon96/dh-go-api-project/internal/domain"
	"github.com/tatianarincon96/dh-go-api-project/internal/product"
	"github.com/tatianarincon96/dh-go-api-project/pkg/store"
	"github.com/tatianarincon96/dh-go-api-project/pkg/web"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type productHandler struct {
	s product.Service
}

// NewProductHandler crea un nuevo controller de productos
func NewProductHandler(s product.Service) *productHandler {
	return &productHandler{
		s: s,
	}
}

func InitProductHandler() *productHandler {
	// Inicializa producto
	list := store.LoadProducts("../products.json")
	repo := product.NewRepository(list)
	service := product.NewService(repo)
	return NewProductHandler(service)
}

// GetAll obtiene todos los productos
func (h *productHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, _ := h.s.GetAll()
		web.Success(ctx, 200, products)
	}
}

// GetByID obtiene un producto por su id
func (h *productHandler) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid id"))
			return
		}
		p, err := h.s.GetById(id)
		if err != nil {
			web.Failure(ctx, 404, errors.New("product not found"))
			return
		}
		web.Success(ctx, 200, p)
	}
}

// Search busca un producto por precio mayor a un valor
func (h *productHandler) Search() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		priceParam := ctx.Query("priceGt")
		price, err := strconv.ParseFloat(priceParam, 64)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid price"))
			return
		}
		products, err := h.s.SearchPriceGt(price)
		if err != nil {
			web.Failure(ctx, 404, errors.New("no products found"))
			return
		}
		web.Success(ctx, 200, products)
	}
}

// ConsumerPrice calcula el precio de un producto
func (h *productHandler) ConsumerPrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ids := ctx.Query("list")
		idsList := strings.Split(ids, ",")
		var intIdList []int
		for _, id := range idsList {
			intId, err := strconv.Atoi(id)
			if err != nil {
				web.Failure(ctx, 400, errors.New("invalid id"))
				return
			}
			intIdList = append(intIdList, intId)
		}

		finalPurchase, err := h.s.GetFinalPurchase(intIdList)
		if err != nil {
			web.Failure(ctx, 404, errors.New("no products found"))
			return
		}
		web.Success(ctx, 200, finalPurchase)
	}
}

// Post crear un producto nuevo
func (h *productHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Valido TOKEN
		token := ctx.GetHeader("TOKEN")
		if token == "" {
			web.Failure(ctx, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(ctx, 401, errors.New("invalid token"))
			return
		}

		// Sigo con la ejecución - Validación de token ok!
		var newProduct domain.Product
		err := ctx.ShouldBindJSON(&newProduct)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid product"))
			return
		}

		valid, err := product.ValidateEmptys(&newProduct)
		if !valid {
			web.Failure(ctx, 400, errors.New(err.Error()))
			return
		}
		valid, err = product.ValidateExpiration(&newProduct)
		if !valid {
			web.Failure(ctx, 400, errors.New(err.Error()))
			return
		}
		p, err := h.s.Create(newProduct)
		if err != nil {
			web.Failure(ctx, 400, errors.New(err.Error()))
			return
		}
		web.Success(ctx, 201, p)
	}
}

// Post crear un producto nuevo
func (h *productHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Valido TOKEN
		token := ctx.GetHeader("TOKEN")
		if token == "" {
			web.Failure(ctx, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(ctx, 401, errors.New("invalid token"))
			return
		}

		var newProduct domain.Product
		err := ctx.ShouldBindJSON(&newProduct)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid product"))
			return
		}
		valid, err := product.ValidateEmptys(&newProduct)
		if !valid {
			web.Failure(ctx, 400, errors.New(err.Error()))
			return
		}
		valid, err = product.ValidateExpiration(&newProduct)
		if !valid {
			web.Failure(ctx, 400, errors.New(err.Error()))
			return
		}
		p, err := h.s.Update(newProduct)
		if err != nil {
			web.Failure(ctx, 400, errors.New(err.Error()))
			return
		}
		web.Success(ctx, 201, p)
	}
}

// DeleteById elimina un producto por su id
func (h *productHandler) DeleteByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Valido TOKEN
		token := ctx.GetHeader("TOKEN")
		if token == "" {
			web.Failure(ctx, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(ctx, 401, errors.New("invalid token"))
			return
		}

		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid id"))
			return
		}
		err = h.s.DeleteByID(id)
		if err != nil {
			web.Failure(ctx, 404, errors.New("product not found"))
			return
		}
		web.Success(ctx, 200, "Product deleted")
	}
}

func (h *productHandler) UpdateField() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Valido TOKEN
		token := ctx.GetHeader("TOKEN")
		if token == "" {
			web.Failure(ctx, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(ctx, 401, errors.New("invalid token"))
			return
		}

		productId := ctx.Param("id")
		id, err := strconv.Atoi(productId)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var productRequest domain.Product
		err = ctx.ShouldBindJSON(&productRequest)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		p, er := h.s.UpdateField(productRequest)
		if er != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("product with id %d not found", id)})
			return
		}
		web.Success(ctx, 200, p)
	}
}

/*func SearchProduct(list []domain.Product) gin.HandlerFunc {
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
}*/
