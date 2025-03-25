package middlewares

import (
	"fmt"
	"net/http"

	"github.com/arash2007mahdavi/web-api-1/api/helper"
	"github.com/gin-gonic/gin"
)

func ApiCheck(ctx *gin.Context) {
	apikey := ctx.GetHeader("api-key")
	if apikey == "1234" {
		ctx.Next()
		return
	}
	ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(false, fmt.Errorf("wrong api-key")))
}
