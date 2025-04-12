package handlers

import (
	"fmt"
	"net/http"

	"github.com/arash2007mahdavi/web-api-1/api/helper"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Health godoc
// @Summary Health Check
// @Description Health Check
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} helper.BaseResponse "Success"
// @Router /v1/health/ [get]
func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("working!...", true))
}

func (h *HealthHandler) HealthById(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(fmt.Sprintf("Working with %v!...", id), true))
}

type User struct {
	Userid string `json:"user_id" binding:"required,id"`
	Browser string `json:"browser" binding:"required,alpha"`
	MobileNum string `json:"number" binding:"required,mobile"`
	Password string `json:"password" binding:"required,password"`
}

func (h *HealthHandler) UserAdd(c *gin.Context) {
	user := User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,  helper.GenerateBaseResponseWithValidationError(false, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(user, true))
}