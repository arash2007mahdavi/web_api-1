package handlers

import "github.com/gin-gonic/gin"

type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "test",
	})
}