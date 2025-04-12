package routers

import (
	"github.com/arash2007mahdavi/web-api-1/api/handlers"
	"github.com/arash2007mahdavi/web-api-1/api/middlewares"
	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()
	r.GET("/", handler.Health)
	r.POST("/:id", middlewares.ApiCheck, handler.HealthById)
	r.POST("/user/add", middlewares.ApiCheck, handler.UserAdd)
}