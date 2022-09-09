package store

import (
	"database/sql"
	"github.com/tatianarincon96/dh-go-api-project/internal/domain"
	"log"
)

type SqlStore struct {
	DB *sql.DB
}

var (
	StorageDB *sql.DB
)

func init() {
	dataSource := "user1@localhost:3306/my_db"
	var err error
	StorageDB, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err = StorageDB.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database connection successful")
}

func (s *SqlStore) GetAll() ([]domain.Product, error) {
	var products []domain.Product
	rows, err := StorageDB.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p domain.Product
		err = rows.Scan(&p.Id, &p.Name, &p.Price, &p.CodeValue)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
