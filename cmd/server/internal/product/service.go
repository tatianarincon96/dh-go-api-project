package product

import (
	"fmt"
	"github.com/tatianarincon96/dh-go-api-project/cmd/server/internal/domain"
)

func AddProduct(list []domain.Product, p domain.Product) []domain.Product {
	list = append(list, p)
	return list
}

func GetProductById(list []domain.Product, id int) (domain.Product, error) {
	for _, product := range list {
		if product.Id == id {
			return product, nil
		}
	}
	return domain.Product{}, fmt.Errorf("Product not found")
}

func GetProductsByQuantity(list []domain.Product, minLimit, maxLimit int) []domain.Product {
	var filteredProducts []domain.Product
	for _, product := range list {
		if product.Quantity >= minLimit && product.Quantity <= maxLimit {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return filteredProducts
}

func BuyProduct(list []domain.Product, codeValue string, quantityToBuy int) domain.Compra {
	var compra domain.Compra
	for _, product := range list {
		if product.CodeValue == codeValue {
			compra.ProductName = product.Name
			compra.Quantity = quantityToBuy
			compra.TotalPrice = product.Price * float64(quantityToBuy)
		}
	}
	return compra
}
