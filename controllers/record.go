package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lzyyauto/auto4play/models"
	"github.com/lzyyauto/auto4play/services"
)

func RecordCommit(c *gin.Context) {
	var commitRecord models.Record
	err := c.BindJSON(&commitRecord)
	if err != nil {
		fmt.Print("提交参数解析错误")
	}
	errs := services.RecordCommit(&commitRecord)
	if len(errs) > 0 {
		c.JSON(200, gin.H{"errcode": 400, "description:": fmt.Sprintf("%v", errs)})
		return
	} else {
		c.JSON(200, gin.H{
			"message": "已提交",
		})
		return
	}
}
