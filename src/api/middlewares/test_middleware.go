package middlewares

import "github.com/gin-gonic/gin"

func TestMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apikey := ctx.GetHeader("api-key")
		if apikey == "1234" {
			ctx.Next()
			return
		}
		ctx.AbortWithStatusJSON(401, gin.H{
			"status": "wrong api-key",
		})
	}
}

func ApiCheck(ctx *gin.Context) {
	apikey := ctx.GetHeader("api-key")
	if apikey == "1234" {
		ctx.Next()
		return
	}
	ctx.AbortWithStatusJSON(401, gin.H{
		"status": "wrong api-key",
	})
}
