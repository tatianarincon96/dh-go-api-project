package products

import "fmt"

type Product struct {
	Id          int
	Name        string
	Quantity    int
	CodeValue   string
	IsPublished bool
	Expiration  string
	Price       float64
}

type Compra struct {
	ProductName string
	Quantity    int
	TotalPrice  float64
}

func AddProduct(p Product) []Product {
	products := getProductList()
	products = append(products, p)
	return products
}

func getProductList() []Product {
	return []Product{
		{Id: 1, Name: "Product 1", Quantity: 100, CodeValue: "1", IsPublished: true, Expiration: "2020-01-01", Price: 1.0},
		{Id: 2, Name: "Product 2", Quantity: 200, CodeValue: "2", IsPublished: true, Expiration: "2020-01-02", Price: 2.0},
		{Id: 3, Name: "Product 3", Quantity: 350, CodeValue: "3", IsPublished: false, Expiration: "2020-01-03", Price: 3.0},
		{Id: 4, Name: "Product 4", Quantity: 399, CodeValue: "4", IsPublished: true, Expiration: "2020-01-04", Price: 4.0},
		{Id: 5, Name: "Product 5", Quantity: 510, CodeValue: "5", IsPublished: false, Expiration: "2020-01-05", Price: 5.0},
	}
}

func GetProductById(id int) (Product, error) {
	products := getProductList()
	for _, product := range products {
		if product.Id == id {
			return product, nil
		}
	}
	return Product{}, fmt.Errorf("product not found")
}

func GetProductsByQuantity(minLimit, maxLimit int) []Product {
	products := getProductList()
	var filteredProducts []Product
	for _, product := range products {
		if product.Quantity >= minLimit && product.Quantity <= maxLimit {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return filteredProducts
}

func BuyProduct(codeValue string, quantityToBuy int) Compra {
	products := getProductList()
	var compra Compra
	for _, product := range products {
		if product.CodeValue == codeValue {
			compra.ProductName = product.Name
			compra.Quantity = quantityToBuy
			compra.TotalPrice = product.Price * float64(quantityToBuy)
		}
	}
	return compra
}
