package services

import (
	"mime/multipart"

	"github.com/VadimBoganov/fulgur/internal/config"
	"github.com/VadimBoganov/fulgur/internal/db/repository"
	"github.com/VadimBoganov/fulgur/internal/domain"
	"github.com/joho/godotenv"
)

type ProductItemService struct {
	repo repository.ProductItem
}

func NewProductItemService(repo repository.ProductItem) *ProductItemService {
	return &ProductItemService{
		repo: repo,
	}
}

func (s *ProductItemService) GetAll() ([]domain.ProductItem, error) {
	return s.repo.GetAll()
}

func (s *ProductItemService) GetById(id int) (*domain.ProductItem, error) {
	return s.repo.GetById(id)
}

func (s *ProductItemService) Add(pi domain.ProductItem, header *multipart.FileHeader) (int64, error) {
	config := config.GetConfig()

	err := godotenv.Load(".env")
	if err != nil {
		return 0, err
	}

	if header != nil {
		err = makeFile(config.LocalFilePath, header)
		if err != nil {
			return 0, err
		}

		err = sendToFtp(config, header)
		if err != nil {
			return 0, err
		}

		pi.ImageUrl = config.FtpUrl + header.Filename

	}

	return s.repo.Insert(pi)
}

func (s *ProductItemService) Update(pi domain.ProductItem, header *multipart.FileHeader) error {
	config := config.GetConfig()

	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	if header != nil {
		err = makeFile(config.LocalFilePath, header)
		if err != nil {
			return err
		}

		err = sendToFtp(config, header)
		if err != nil {
			return err
		}

		pi.ImageUrl = config.FtpUrl + header.Filename
	}

	return s.repo.Update(pi)
}

func (s *ProductItemService) Remove(id int) error {
	return s.repo.Remove(id)
}
