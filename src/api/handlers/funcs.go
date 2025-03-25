package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, "Working!...")
}

func (h *HealthHandler) HealthById(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, fmt.Sprintf("Working with %v!...", id))
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
		c.AbortWithStatusJSON(http.StatusBadRequest,  gin.H{
			"status": http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": user,
	})
}