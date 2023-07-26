package services

import (
	"github.com/VadimBoganov/fulgur/internal/db/repository"
	"github.com/VadimBoganov/fulgur/internal/domain"
)

type Product interface {
	GetAll() ([]domain.Product, error)
	Add([]domain.Product) (int64, error)
	UpdateById(id int, name string) error
	RemoveById(id int) error
}

type ProductType interface {
	GetAll() ([]domain.ProductType, error)
	Add(domain.ProductType) (int64, error)
	Update(domain.ProductType) error
	Remove(id int) error
}

type Service struct {
	Product
	ProductType
}

func NewService(productRepo *repository.ProductRepository, productTypeRepo *repository.ProductTypeRepository) *Service {
	return &Service{
		Product:     NewProductService(productRepo),
		ProductType: NewProductTypeService(productTypeRepo),
	}
}
