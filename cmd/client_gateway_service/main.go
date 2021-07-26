package main

import (
	"cafe/pkg/client_gateway_service"
	"fmt"
	yconfig "github.com/rowdyroad/go-yaml-config"
)

func main() {
	var config client_gateway_service.Config
	yconfig.LoadConfig(&config, "configs/client_gateway_service/config.yaml", nil)
	fmt.Printf("%+v\n", config)
	app := client_gateway_service.NewApp(config)
	defer app.Close()
	app.Run()
}
