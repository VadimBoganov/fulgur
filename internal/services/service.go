package services

import (
	"github.com/VadimBoganov/fulgur/internal/db/repository"
	"github.com/VadimBoganov/fulgur/internal/domain"
)

type Product interface {
	GetAll() ([]domain.Product, error)
	Add([]domain.Product) error
	RemoveById(id int) error
}

type Service struct {
	Product
}

func NewService(repo *repository.ProductRepository) *Service {
	return &Service{Product: NewProductService(repo)}
}
