package main

import (
	"github.com/arash2007mahdavi/web-api-1/api"
	"github.com/arash2007mahdavi/web-api-1/config"
	"github.com/arash2007mahdavi/web-api-1/data/cache"
	"github.com/arash2007mahdavi/web-api-1/data/database"
	"github.com/arash2007mahdavi/web-api-1/pkg/logging"
)

func main() {
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)

	cache.InitRedis(cfg)
	defer cache.CloseRedis()

	err := database.InitDb(cfg)
	if err != nil {
		logger.Warn(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	defer database.CloseDb()

	api.InitServer(cfg)
}