package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
	"web-study/controller"
	"web-study/logger"
	"web-study/middlewares"
)

func Setup() *gin.Engine {
	r := gin.New()
	/*
	   解决跨域问题:
	   cors.New方法返回一个函数参数是c *gin.Context
	   将这个参数赋值给mwCORS,直接当中间间使用,
	   默认修改返回的请求头,实现跨域功能
	   cors.Config为一个结构体,结构体实例后传入cors.New实现生成中间件功能
	*/
	mwCORS := cors.New(cors.Config{
		//准许跨域请求网站,多个使用,分开,限制使用*
		AllowOrigins: []string{"*"},
		//准许使用的请求方式
		AllowMethods: []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		//准许使用的请求表头
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
		//显示的请求表头
		ExposeHeaders: []string{"Content-Type"},
		//凭证共享,确定共享
		AllowCredentials: true,
		//容许跨域的原点网站,可以直接return true就万事大吉了
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		//超时时间设定
		MaxAge: 24 * time.Hour,
	})
	r.Use(logger.GinLogger(), logger.GinRecovery(true), mwCORS)

	//测试
	r.GET("/test", middlewares.JWTAuthMiddleware(), func(context *gin.Context) {
		//如果是登录用户，判断请求头中是否有 有效的JWT

		context.JSON(200, gin.H{
			"msg": "测试成功",
		})
	})

	user := r.Group("/user")
	{
		//注册用户
		user.POST("/signup", controller.SignUpUser)
		//登录功能
		user.POST("/login", controller.LoginUp)
	}

	community := r.Group("/comunity")
	{
		//向社区插入数据
		community.POST("/insert", controller.InsertComData)
		//获取所有社区
		community.GET("/list", controller.GetCommunityList)
	}
	return r
}
