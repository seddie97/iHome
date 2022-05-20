package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"iHome/web/controller"
	"iHome/web/model"
)

func LoginFilter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//初始化session对象
		s := sessions.Default(ctx)
		userName := s.Get("userName")
		if userName == nil {
			ctx.Abort() //从这里返回，不必继续执行
		} else {
			ctx.Next()
		}
	}
}

// 添加gin框架开发3步骤
func main() {

	//初始化redis连接池
	model.InitRedis()

	//Mysql连接池
	model.InitDb()

	// 初始化路由
	router := gin.Default()

	//初始化容器
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))

	//使用容器-----使用中间件！指定容器，对以后的路由生效
	router.Use(sessions.Sessions("mysession", store))

	// 路由匹配
	/*	router.GET("/", func(context *gin.Context) {
		context.Writer.WriteString("项目开始了....")
	})*/
	router.Static("/home", "web/view")
	//添加路由分组
	r1 := router.Group("/api/v1.0")
	{
		r1.GET("/session", controller.GetSession)
		r1.GET("/imagecode/:uuid", controller.GetImageCd)
		r1.POST("/users", controller.PostRet)
		r1.GET("/areas", controller.GetArea)
		r1.POST("/sessions", controller.PostLogin)

		r1.Use(LoginFilter()) //以后的路由都不需要校验session了，直接获取数据即可

		r1.DELETE("/session", controller.DeleteSession)
		r1.GET("/user", controller.GetUserInfo)
		r1.PUT("/user/name", controller.PutUserInfo)
		r1.POST("/user/avatar", controller.PostAvatar)
		r1.POST("/user/auth", controller.PostUserAuth)
	}

	// 启动运行
	router.Run(":8080")
}
