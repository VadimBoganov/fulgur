package services

import (
	"io"
	"mime/multipart"
	"os"
	"time"

	"github.com/VadimBoganov/fulgur/internal/config"
	"github.com/VadimBoganov/fulgur/internal/db/repository"
	"github.com/VadimBoganov/fulgur/internal/domain"
	"github.com/jlaffaye/ftp"
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

type ProductSubtype interface {
	GetAll() ([]domain.ProductSubType, error)
	Add(domain.ProductSubType) (int64, error)
	Update(domain.ProductSubType) error
	Remove(id int) error
}

type ProductItem interface {
	GetAll() ([]domain.ProductItem, error)
	GetById(id int) (*domain.ProductItem, error)
	Add(domain.ProductItem, *multipart.FileHeader) (int64, error)
	Update(domain.ProductItem, *multipart.FileHeader) error
	Remove(id int) error
}

type Item interface {
	GetAll() ([]domain.Item, error)
	GetById(id int) (*domain.Item, error)
	Add(domain.Item, *multipart.FileHeader) (int64, error)
	Update(domain.Item, *multipart.FileHeader) error
	Remove(id int) error
}

type User interface {
	GetAll() ([]domain.User, error)
	GetById(id int) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	Add(*domain.User) (int64, error)
	Update(*domain.User) error
	Remove(id int) error
}

type Service struct {
	Product
	ProductType
	ProductSubtype
	ProductItem
	Item
	User
}

func NewService(productRepo *repository.ProductRepository, productTypeRepo *repository.ProductTypeRepository, productSubtypeRepo *repository.ProductSubtypeRepository, productItemRepo *repository.ProductItemRepository, ItemRepo *repository.ItemRepository, UserRepo *repository.UserRepository) *Service {
	return &Service{
		Product:        NewProductService(productRepo),
		ProductType:    NewProductTypeService(productTypeRepo),
		ProductSubtype: NewProductSubtypeService(productSubtypeRepo),
		ProductItem:    NewProductItemService(productItemRepo),
		Item:           NewItemService(ItemRepo),
		User:           NewUserService(UserRepo),
	}
}

func makeFile(filePath string, header *multipart.FileHeader) error {
	infile, err := header.Open()
	if err != nil {
		return err
	}

	defer infile.Close()

	var outfile *os.File
	if outfile, err = os.Create(filePath + header.Filename); nil != err {
		return err
	}

	if _, err = io.Copy(outfile, infile); nil != err {
		return err
	}

	return nil
}

func sendToFtp(config *config.Config, header *multipart.FileHeader) error {
	c, err := ftp.Dial(os.Getenv("FTP_HOST")+":"+os.Getenv("FTP_PORT"), ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		return err
	}

	err = c.Login(os.Getenv("FTP_LOGIN"), os.Getenv("FTP_PASS"))
	if err != nil {
		return err
	}

	fileName := header.Filename
	file, err := os.Open(config.LocalFilePath + fileName)
	if err != nil {
		return err
	}

	err = c.Stor(config.FTPPath+fileName, file)
	if err != nil {
		return err
	}

	if err = c.Quit(); err != nil {
		return err
	}

	return nil
}
