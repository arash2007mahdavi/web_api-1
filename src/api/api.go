package api

import (
	"fmt"

	"github.com/arash2007mahdavi/golang-web-api-1/api/routers"
	"github.com/arash2007mahdavi/golang-web-api-1/config"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1/")
	{
		health := v1.Group("/health")
		routers.Health(health)
	}

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}
