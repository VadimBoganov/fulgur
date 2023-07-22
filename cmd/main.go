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

	var config = &config2.Config{}

	if err := config.InitConfig("internal/config"); err != nil {
		logger.Error("Error occurred while initialize config: %s", err.Error())
	}

	db := db2.NewDB(config.DatabasePath)
	if db == nil {
		logger.Error("Cant open db...")
	}

	db2.RunMigrations(db)

	productRepository := repository2.NewProductRespository(db)
	service := services.NewService(productRepository)
	handler := handlers.NewHandler(service)

	server := new(fulgur.Server)
	if err := server.Run(config.Server.Port, handler.InitRoutes()); err != nil {
		logger.Fatalf("Error occured while runnig http server: %s", err.Error())
	}
}
