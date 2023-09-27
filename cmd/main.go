package main

import (
	fulgur "github.com/VadimBoganov/fulgur/internal"
	config2 "github.com/VadimBoganov/fulgur/internal/config"
	db2 "github.com/VadimBoganov/fulgur/internal/db"
	repository2 "github.com/VadimBoganov/fulgur/internal/db/repository"
	"github.com/VadimBoganov/fulgur/internal/handlers"
	"github.com/VadimBoganov/fulgur/internal/services"
	logger2 "github.com/VadimBoganov/fulgur/pkg/logging"
)

func main() {
	logger := logger2.GetLogger()

	if err := config2.InitConfig("internal/config"); err != nil {
		logger.Errorf("Error occurred while initialize config: %s", err.Error())
	}

	config := config2.GetConfig()

	db := db2.NewDB(config.DatabasePath)
	if db == nil {
		logger.Error("Cant open db...")
	}

	db2.RunMigrations(db)

	productRepository := repository2.NewProductRespository(db)
	productTypeRepo := repository2.NewProductTypeRepository(db)
	productSubtypeRepo := repository2.NewProductSubtypeRepository(db)
	productItemRepo := repository2.NewProductItemRepository(db)
	itemRepo := repository2.NewItemRepository(db)
	service := services.NewService(productRepository, productTypeRepo, productSubtypeRepo, productItemRepo, itemRepo)
	handler := handlers.NewHandler(service)

	server := new(fulgur.Server)
	if err := server.Run(config.Server.Port, handler.InitRoutes()); err != nil {
		logger.Fatalf("Error occured while runnig http server: %s", err.Error())
	}
}
