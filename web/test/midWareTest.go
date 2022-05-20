package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

//创建中间件
func Test1(ctx *gin.Context) {
	fmt.Println("1111111")
	//return
	t := time.Now()
	ctx.Next()
	fmt.Println(time.Now().Sub(t))
	//ctx.Abort()
}

//创建另一种格式的中间件
func Test2() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("333333333")
		context.Abort()
		fmt.Println("555555555")
	}
}

func main() {
	router := gin.Default()

	//使用中间件
	router.Use(Test1)
	router.Use(Test2())

	router.GET("/midWareTest", func(context *gin.Context) {
		fmt.Println("22222222")
		context.Writer.WriteString("ldj最帅！")
	})

	router.Run(":1231")
}
