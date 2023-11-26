package services

import (
	"github.com/VadimBoganov/fulgur/internal/db/repository"
	"github.com/VadimBoganov/fulgur/internal/domain"
)

type ProductTypeService struct {
	repo repository.ProductType
}

func NewProductTypeService(repo repository.ProductType) *ProductTypeService {
	return &ProductTypeService{
		repo: repo,
	}
}

func (s *ProductTypeService) GetAll() ([]domain.ProductType, error) {
	return s.repo.GetAll()
}

func (s *ProductTypeService) GetByProductId(productId int) ([]domain.ProductType, error) {
	return s.repo.GetByProductId(productId)
}

func (s *ProductTypeService) Add(pt domain.ProductType) (int64, error) {
	return s.repo.Insert(pt)
}

func (s *ProductTypeService) Update(pt domain.ProductType) error {
	return s.repo.Update(pt)
}

func (s *ProductTypeService) Remove(id int) error {
	return s.repo.Remove(id)
}
