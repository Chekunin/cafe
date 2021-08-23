package main

import (
	"cafe/pkg/review_media_storage_service"
	yconfig "github.com/rowdyroad/go-yaml-config"
)

func main() {
	var config review_media_storage_service.Config
	yconfig.LoadConfig(&config, "configs/nsi_service/config.yaml", nil)
	app := review_media_storage_service.NewApp(config)
	defer app.Close()
	app.Run()
}
