package services

import (
	"github.com/VadimBoganov/fulgur/internal/db/repository"
	"github.com/VadimBoganov/fulgur/internal/domain"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetAll() ([]domain.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetById(id int) (*domain.User, error) {
	return s.repo.GetById(id)
}

func (s *UserService) GetByEmail(email string) (*domain.User, error) {
	return s.repo.GetByEmail(email)
}

func (s *UserService) Add(user *domain.User) (int64, error) {
	return s.repo.Insert(user)
}

func (s *UserService) Update(user *domain.User) error {
	return s.repo.Update(user)
}

func (s *UserService) Remove(id int) error {
	return s.repo.Remove(id)
}
