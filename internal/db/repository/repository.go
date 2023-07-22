package repository

import (
	"database/sql"

	"github.com/VadimBoganov/fulgur/internal/domain"
)

type Product interface {
	GetAll() ([]domain.Product, error)
	GetAllTypes() ([]domain.ProductType, error)
	GetAllSubTypes() ([]domain.ProductSubType, error)

	Insert(products []domain.Product) error

	UpdateById(id int, name string) error

	RemoveById(id int) error
	RemoveAll() error
}

type ProductItem interface {
	GetAll() ([]domain.ProductItem, error)
	Insert(productItems []domain.ProductItem) error
	Remove(productItems []domain.ProductItem) error
}

type Item interface {
	GetAll() ([]domain.Item, error)
	Insert(items []domain.Item) error
	Remove(items []domain.Item) error
}

type Repository struct {
	Product
	ProductItem
	Item
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Product: NewProductRespository(db),
	}
}
