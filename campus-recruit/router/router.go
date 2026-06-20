package router

import (
	"campus-recruit/controller"
	"campus-recruit/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 全局中间件
	r.Use(middleware.CORS())

	// 健康检查
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "pong"})
	})

	// API 路由组
	api := r.Group("/api/v1")

	// 无需登录的接口
	{
		api.POST("/register", controller.Register)
		api.POST("/login", controller.Login)
		// TODO: 职位列表（无需登录可查看）
		// api.GET("/positions", controller.GetPositionList)
		// api.GET("/positions/:id", controller.GetPositionDetail)
	}

	// 需要登录的接口
	auth := api.Group("")
	auth.Use(middleware.JWTAuth())
	{
		auth.GET("/user/info", controller.GetUserInfo)

		// TODO: 企业端接口
		// auth.POST("/positions", controller.CreatePosition)
		// auth.PUT("/positions/:id", controller.UpdatePosition)

		// TODO: 学生端接口
		// auth.POST("/resume", controller.UploadResume)
		// auth.POST("/apply/:position_id", controller.Apply)
		// auth.POST("/favorite/:position_id", controller.Favorite)
	}

	return r
}
