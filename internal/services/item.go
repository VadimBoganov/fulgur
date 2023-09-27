package services

import (
	"mime/multipart"

	"github.com/VadimBoganov/fulgur/internal/config"
	"github.com/VadimBoganov/fulgur/internal/db/repository"
	"github.com/VadimBoganov/fulgur/internal/domain"
	"github.com/joho/godotenv"
)

type ItemService struct {
	repo repository.Item
}

func NewItemService(repo repository.Item) *ItemService {
	return &ItemService{
		repo: repo,
	}
}

func (s *ItemService) GetAll() ([]domain.Item, error) {
	return s.repo.GetAll()
}

func (s *ItemService) Add(item domain.Item, header *multipart.FileHeader) (int64, error) {
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

		item.ImageUrl = config.FtpUrl + header.Filename
	}

	return s.repo.Insert(item)
}

func (s *ItemService) Update(item domain.Item, header *multipart.FileHeader) error {
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

		item.ImageUrl = config.FtpUrl + header.Filename
	}

	return s.repo.Update(item)
}

func (s *ItemService) Remove(id int) error {
	return s.repo.Remove(id)
}
