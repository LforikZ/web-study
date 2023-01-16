package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"web-study/controller"
	"web-study/logger"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//测试
	r.GET("/test", func(context *gin.Context) {
		fmt.Println("test启动")
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
