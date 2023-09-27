package services_test

import (
	"os"
	"testing"

	"github.com/VadimBoganov/fulgur/internal/domain"
	"github.com/VadimBoganov/fulgur/internal/services"
)

var service *services.ProductItemService

func TestMain(m *testing.M) {
	service = services.NewProductItemService(nil)
	os.Exit(m.Run())
}

func TestProductItem_Add_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	productItem := domain.ProductItem{
		Id:               1,
		ProductSubTypeId: 1,
		Name:             "name",
		ImageUrl:         "url",
	}

	service.Add(productItem, nil)
}
