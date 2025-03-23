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

func (h *HealthHandler) HeaderBinder1(c *gin.Context) {
	userid := c.GetHeader("Userid")
	c.JSON(200, gin.H{
		"result": "HeaderBinder1",
		"userid": userid,
	})
}

type header22 struct {
	Userid string `json:"user_id" binding:"required,email"`
	Browser string `json:"browser" binding:"required,alpha"`
	MobileNum string `json:"number" binding:"required,mobile"`
	Password string `json:"password" binding:"required,password"`
}

func (h *HealthHandler) HeaderBinder2(c *gin.Context) {
	hea := header22{}
	err := c.ShouldBindJSON(&hea)
	if err != nil {
		c.AbortWithStatusJSON(403,  gin.H{
			"status": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"result": "HeaderBinder1",
		"header": hea,
	})
}