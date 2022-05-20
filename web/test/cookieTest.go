package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	//初始化路由
	router := gin.Default()

	//初始化容器
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	//store.Options(sessions.Options{
	//	MaxAge: 0,
	//})
	//使用容器
	router.Use(sessions.Sessions("mysession", store))

	//路由匹配
	router.GET("/test", func(context *gin.Context) {
		//设置cookie
		//func (c *Context) SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)
		//context.SetCookie("mytest", "chuanzhi", 0, "", "", false, true)

		//获取cookie
		//str, err := context.Cookie("mytest")
		//if err != nil {
		//    fmt.Println("Cookie err:", err)
		//    return
		//}

		//fmt.Println(str)

		//调用session，设置session数据
		s := sessions.Default(context)
		//设置session
		//s.Set("ldj", "zuishuai!")
		//修改session时，需要Save函数配合，不然不生效
		//s.Save()
		v := s.Get("ldj")
		fmt.Println(v.(string))

		context.Writer.WriteString("测试session")
	})

	//启动
	router.Run(":9999")
}
