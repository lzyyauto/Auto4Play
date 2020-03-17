package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertRecord(c *gin.Context) {
	thing := c.Param("thing")
	fmt.Printf("thing:%s", thing)

	c.String(http.StatusOK, thing)
}

func GetTestGin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "oh",
	})
}
