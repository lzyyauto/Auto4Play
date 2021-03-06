package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lzyyauto/auto4play/controllers"
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

		v1.GET("/record/quick/:thing", controllers.QuickRecordCommit)

		v1.POST("/record/commit", controllers.RecordCommit)
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
