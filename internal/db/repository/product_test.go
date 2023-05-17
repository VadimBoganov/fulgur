package repository_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	logger2 "github.com/VadimBoganov/fulgur/internal/logging"

	config2 "github.com/VadimBoganov/fulgur/internal/config"

	"github.com/VadimBoganov/fulgur/internal/db"

	"github.com/VadimBoganov/fulgur/internal/db/repository"
	"github.com/VadimBoganov/fulgur/internal/domain"
)

var logger = logger2.GetLogger()
var products []domain.Product
var repo *repository.ProductRepository
var config = &config2.Config{}

func TestMain(m *testing.M) {
	products = []domain.Product{{Name: "pripoi"}, {Name: "metal"}}

	if err := config.InitConfig("../../config"); err != nil {
		logger.Error("Error occured while initialize config: ", err.Error())
	}

	database := db.NewDB(config.DatabasePath)
	repo = repository.NewProductRespository(database)

	os.Exit(m.Run())
}

func TestProductRepository_Insert_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	_ = repo.Insert(products)

	data, _ := repo.GetAll()

	assert.NotNil(t, data)
	assert.Greater(t, len(data), 0)

	_ = repo.RemoveAll()
}

func TestProductRepository_RemoveById_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	_ = repo.Insert(products)
	_ = repo.RemoveById(1)

	data, _ := repo.GetAll()

	assert.Equal(t, 1, len(data))
	assert.Equal(t, 2, data[0].Id)

	_ = repo.RemoveAll()
}
