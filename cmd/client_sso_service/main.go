package main

import (
	"cafe/pkg/client_sso_service"
	"fmt"
	yconfig "github.com/rowdyroad/go-yaml-config"
)

func main() {
	var config client_sso_service.Config
	yconfig.LoadConfig(&config, "configs/client_sso_service/config.yaml", nil)
	fmt.Printf("%+v\n", config)
	app := client_sso_service.NewApp(config)
	defer app.Close()
	app.Run()
}
