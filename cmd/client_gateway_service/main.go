package main

import (
	_ "cafe/doc/client_gateway_service/v1"
	"cafe/pkg/client_gateway_service"
	yconfig "github.com/rowdyroad/go-yaml-config"
)

// @title Client gateway service
// @version 1.0
// @description Gateway service for simple users
// @BasePath /client-gateway/v1

func main() {
	var config client_gateway_service.Config
	yconfig.LoadConfig(&config, "configs/client_gateway_service/config.yaml", nil)
	app := client_gateway_service.NewApp(config)
	defer app.Close()
	app.Run()
}
