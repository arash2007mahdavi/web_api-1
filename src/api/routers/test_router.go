package routers

import (
	"github.com/arash2007mahdavi/golang-web-api-1/api/handlers"
	"github.com/gin-gonic/gin"
)

func TestRouter(r *gin.RouterGroup) {
	h := handlers.NewTestHandler()
	r.GET("/", h.Test)
}