package services

import (
	"github.com/VadimBoganov/fulgur/internal/db/repository"
	"github.com/VadimBoganov/fulgur/internal/domain"
)

type ProductService struct {
	repo repository.Product
}

func NewProductService(repo repository.Product) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) GetAll() ([]domain.Product, error) {
	return s.repo.GetAll()
}

func (s *ProductService) Add(products []domain.Product) error {
	return s.repo.Insert(products)
}

func (s *ProductService) UpdateById(id int, name string) error {
	return s.repo.UpdateById(id, name)
}

func (s *ProductService) RemoveById(id int) error {
	return s.repo.RemoveById(id)
}
