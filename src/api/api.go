package api

import (
	"fmt"

	"github.com/arash2007mahdavi/golang-web-api-1/api/routers"
	"github.com/arash2007mahdavi/golang-web-api-1/api/validations"
	"github.com/arash2007mahdavi/golang-web-api-1/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()

	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validations.IranainMobileNumberValidator, true)
	}

	val2, ok2 := binding.Validator.Engine().(*validator.Validate)
	if ok2 {
		val2.RegisterValidation("password", validations.PasswoordValidator, true)
	}

	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1/")
	{
		health := v1.Group("/health")
		routers.Health(health)
	}

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}
