package main

import (
	"cafe/pkg/nsi_service"
	yconfig "github.com/rowdyroad/go-yaml-config"
)

func main() {
	var config nsi_service.Config
	yconfig.LoadConfig(&config, "configs/nsi_service/config.yaml", nil)
	app := nsi_service.NewApp(config)
	defer app.Close()
	app.Run()
}
