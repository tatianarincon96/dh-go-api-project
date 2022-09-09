package product

import (
	"errors"
	"github.com/tatianarincon96/dh-go-api-project/internal/domain"
	"math"
)

type Service interface {
	GetAll() ([]domain.Product, error)
	GetById(id int) (domain.Product, error)
	SearchPriceGt(price float64) ([]domain.Product, error)
	GetFinalPurchase(intIdList []int) (domain.Compra, error)
	Create(p domain.Product) (domain.Product, error)
	Update(p domain.Product) (domain.Product, error)
	DeleteByID(id int) error
	UpdateField(p domain.Product) (domain.Product, error)
}

type service struct {
	r Repository
}

// NewSerice crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

// GetAll devuelve todos los productos
func (s *service) GetAll() ([]domain.Product, error) {
	return s.r.GetAll(), nil
}

// GetByID busca un producto por id
func (s *service) GetById(id int) (domain.Product, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

// SearchPriceGt busca productos por precio mayor que el precio dado
func (s *service) SearchPriceGt(price float64) ([]domain.Product, error) {
	l := s.r.SearchPriceGt(price)
	if len(l) == 0 {
		return []domain.Product{}, errors.New("no products found")
	}
	return l, nil
}

// GetFinalPurchase devuelve el total de la compra
func (s *service) GetFinalPurchase(intIdList []int) (domain.Compra, error) {
	var products []domain.Product
	var totalPrice float64
	for _, id := range intIdList {
		product, err := s.r.GetByID(id)
		if err != nil {
			return domain.Compra{}, err
		}
		if product.IsPublished && product.Quantity > 0 {
			products = append(products, product)
			totalPrice += product.Price
		}
	}
	if len(products) < 10 {
		totalPrice *= 1.21
	} else if len(products) <= 20 {
		totalPrice *= 1.17
	} else {
		totalPrice *= 1.15
	}
	return domain.Compra{Products: products, TotalPrice: math.Round(totalPrice*100) / 100}, nil
}

// Create agrega un nuevo producto
func (s *service) Create(p domain.Product) (domain.Product, error) {
	p, err := s.r.Create(p)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

// Update actualiza un nuevo producto
func (s *service) Update(p domain.Product) (domain.Product, error) {
	p, err := s.r.Update(p)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

// DeleteByID elimina un producto por id
func (s *service) DeleteByID(id int) error {
	err := s.r.DeleteByID(id)
	if err != nil {
		return err
	}
	return nil
}

// Patch actualiza un campo del producto
func (s *service) UpdateField(p domain.Product) (domain.Product, error) {
	product, err := s.r.UpdateField(p.Id, p)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

/*
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
}*/
