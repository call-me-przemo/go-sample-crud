package controllers

import (
	"net/http"

	"github.com/call-me-przemo/go-sample-crud/services"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	msg := services.Dumb()

	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}
