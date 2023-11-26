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
	GetByProductId(productId int) ([]domain.ProductType, error)
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
	GetById(id int) (*domain.ProductItem, error)
	Insert(pi domain.ProductItem) (int64, error)
	Update(pi domain.ProductItem) error
	Remove(id int) error
}

type Item interface {
	GetAll() ([]domain.Item, error)
	GetByProductItemId(productItemId int) ([]domain.Item, error)
	GetById(id int) (*domain.Item, error)
	Insert(item domain.Item) (int64, error)
	Update(item domain.Item) error
	Remove(id int) error
}

type User interface {
	GetAll() ([]domain.User, error)
	GetById(id int) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	Insert(item *domain.User) (int64, error)
	Update(item *domain.User) error
	Remove(id int) error
}

type Repository struct {
	Product
	ProductType
	ProductSubtype
	ProductItem
	Item
	User
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Product:        NewProductRespository(db),
		ProductType:    NewProductTypeRepository(db),
		ProductSubtype: NewProductSubtypeRepository(db),
		ProductItem:    NewProductItemRepository(db),
		Item:           NewItemRepository(db),
		User:           NewUserRepository(db),
	}
}
