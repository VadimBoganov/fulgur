package repository

import (
	"database/sql"

	"github.com/VadimBoganov/fulgur/internal/domain"
)

type Product interface {
	GetAll() ([]domain.Product, error)

	// TODO: rewrite to InsertAll(InsertBulk) and add insert product
	Insert(products []domain.Product) (int64, error)

	// TODO: name update and rewrite signature
	UpdateById(id int, name string) error

	Remove(id int) error
	RemoveAll() error
}

type ProductType interface {
	GetAll() ([]domain.ProductType, error)
	Insert(pt domain.ProductType) (int64, error)
	Update(pt domain.ProductType) error
	Remove(id int) error
}

type ProductSubtype interface {
	GetAll() ([]domain.ProductSubType, error)
	Insert(pst domain.ProductSubType) (int64, error)
	Update(pst domain.ProductSubType) error
	Remove(id int) error
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
	ProductType
	ProductSubtype
	ProductItem
	Item
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Product:        NewProductRespository(db),
		ProductType:    NewProductTypeRepository(db),
		ProductSubtype: NewProductSubtypeRepository(db),
	}
}
