package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lzyyauto/auto4play/services"
)

// Init 初始化路由godef -debug -f main.go net.ResolveTCPAddr
func Init() error {
	router := gin.Default()
	// CrossDomain跨域处理，options请求处理
	// router.Use(middleware.CrossDomain())
	// v1群组对任何人开放
	v1 := router.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		v1.GET("/helloworld", HelloWordGet)
		v1.GET("/record/:thing", services.InsertRecord)
		v1.GET("/pong", services.GetTestGin)
	}

	// v2 := router.Group("/v2")
	// v2群组使用中间件AuthMiddleWare，需要token权限才能请求到
	// v2.Use(middleware.AuthMiddleWare())
	// {
	// 	v2.POST("/publish", post.Publish)
	// 	v2.POST("/isload", user.IsLoad)
	// 	v2.POST("/reply1", post.Reply1)
	// 	v2.POST("/reply2", post.Reply2)
	// }
	return router.Run(":8000")
}

func HelloWordGet(c *gin.Context) {
	c.String(http.StatusOK, "hello world get")
}
