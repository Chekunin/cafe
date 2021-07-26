package main

import (
	"cafe/pkg/nsi_service"
	"fmt"
	yconfig "github.com/rowdyroad/go-yaml-config"
)

func main() {
	var config nsi_service.Config
	yconfig.LoadConfig(&config, "configs/nsi_service/config.yaml", nil)
	fmt.Printf("%+v\n", config)
	app := nsi_service.NewApp(config)
	defer app.Close()
	app.Run()
}
