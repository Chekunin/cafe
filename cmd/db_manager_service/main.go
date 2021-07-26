package main

import (
	"cafe/pkg/db_manager_service"
	"fmt"
	yconfig "github.com/rowdyroad/go-yaml-config"
)

func main() {
	var config db_manager_service.Config
	yconfig.LoadConfig(&config, "configs/db_manager_service/config.yaml", nil)
	fmt.Printf("%+v\n", config)
	app := db_manager_service.NewApp(config)
	defer app.Close()
	app.Run()
}
