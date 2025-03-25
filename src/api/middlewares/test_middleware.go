package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiCheck(ctx *gin.Context) {
	apikey := ctx.GetHeader("api-key")
	if apikey == "1234" {
		ctx.Next()
		return
	}
	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"status": http.StatusBadRequest,
		"message": "wrong api-key",
	})
}
