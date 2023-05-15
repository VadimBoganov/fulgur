package main

import (
	config2 "github.com/VadimBoganov/fulgur/internal/config"
	db2 "github.com/VadimBoganov/fulgur/internal/db"
	logger2 "github.com/VadimBoganov/fulgur/internal/logging"
)

func main() {
	logger := logger2.GetLogger()

	var config = &config2.Config{}

	if err := config.InitConfig("internal/config"); err != nil {
		logger.Error("Error occured while initialize config: %s", err.Error())
	}

	db := db2.NewDB(config.DatabasePath)
	if db == nil {
		logger.Error("Cant open db...")
	}

    db2.RunMigrations(db)
}
