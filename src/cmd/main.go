package main

import (
	"github.com/arash2007mahdavi/web-api-1/api"
	"github.com/arash2007mahdavi/web-api-1/config"
	"github.com/arash2007mahdavi/web-api-1/data/cache"
)

func main() {
	cfg := config.GetConfig()
	cache.InitRedis(cfg)
	defer cache.CloseRedis()
	api.InitServer(cfg)
}