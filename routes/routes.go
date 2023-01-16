package routes

import (
	"github.com/gin-gonic/gin"
	"web-study/controller"
	"web-study/logger"
	"web-study/middlewares"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//测试
	r.GET("/test", middlewares.JWTAuthMiddleware(), func(context *gin.Context) {
		//如果是登录用户，判断请求头中是否有 有效的JWT

		context.JSON(200, gin.H{
			"msg": "测试成功",
		})
	})

	//注册用户
	r.POST("/signup", controller.SignUpUser)

	//登录功能
	r.POST("/login", controller.LoginUp)

	return r
}
