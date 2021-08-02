package main

import (
	"cafe/pkg/db_manager_service"
	yconfig "github.com/rowdyroad/go-yaml-config"
)

func main() {
	var config db_manager_service.Config
	yconfig.LoadConfig(&config, "configs/db_manager_service/config.yaml", nil)
	app := db_manager_service.NewApp(config)
	defer app.Close()
	app.Run()
}
