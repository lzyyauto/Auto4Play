package controllers

import (
	"github.com/lzyyauto/auto4play/util"

	"github.com/gin-gonic/gin"
	"github.com/lzyyauto/auto4play/config"
)

func RemindThing(c *gin.Context) {
	thing := c.Param("thing")
	serviceCfg := config.ServiceCfg
	url := serviceCfg.Host + ":" + serviceCfg.Port + "/v1" + "/record/quick/" + thing
	err := util.SendMessage(url)
	if err == nil {
		c.JSON(200, gin.H{
			"message": "已提醒",
		})
		return
	}
}
