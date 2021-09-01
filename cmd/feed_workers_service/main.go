package main

import (
	"cafe/pkg/feed_workers_service"
	yconfig "github.com/rowdyroad/go-yaml-config"
)

func main() {
	var config feed_workers_service.Config
	yconfig.LoadConfig(&config, "configs/feed_workers_service/config.yaml", nil)
	app := feed_workers_service.NewApp(config)
	defer app.Close()
	app.Run()
}
