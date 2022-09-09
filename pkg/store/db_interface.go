package store

import "github.com/tatianarincon96/dh-go-api-project/internal/domain"

type StoreInterface interface {
	GetAll() ([]domain.Product, error)
}
