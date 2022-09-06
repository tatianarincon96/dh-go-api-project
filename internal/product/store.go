package product

import (
	"encoding/json"
	"github.com/tatianarincon96/dh-go-api-project/internal/domain"
	"os"
)

func LoadProducts(path string) []domain.Product {
	list := []domain.Product{}

	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(file, &list)
	if err != nil {
		panic(err)
	}
	return list
}
