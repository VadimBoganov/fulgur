package services

import (
	"github.com/VadimBoganov/fulgur/internal/db/repository"
	"github.com/VadimBoganov/fulgur/internal/domain"
)

type ProductSubtypeService struct {
	repo repository.ProductSubtype
}

func NewProductSubtypeService(repo repository.ProductSubtype) *ProductSubtypeService {
	return &ProductSubtypeService{
		repo: repo,
	}
}

func (s *ProductSubtypeService) GetAll() ([]domain.ProductSubType, error) {
	return s.repo.GetAll()
}

func (s *ProductSubtypeService) Add(pst domain.ProductSubType) (int64, error) {
	return s.repo.Insert(pst)
}

func (s *ProductSubtypeService) Update(pst domain.ProductSubType) error {
	return s.repo.Update(pst)
}

func (s *ProductSubtypeService) Remove(id int) error {
	return s.repo.Remove(id)
}
