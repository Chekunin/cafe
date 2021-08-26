package main

import (
	"cafe/pkg/feed_workpool_service"
	yconfig "github.com/rowdyroad/go-yaml-config"
)

func main() {
	var config feed_workpool_service.Config
	yconfig.LoadConfig(&config, "configs/feed_workpool_service/config.yaml", nil)
	app := feed_workpool_service.NewApp(config)
	defer app.Close()
	app.Run()
}
