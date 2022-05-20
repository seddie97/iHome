package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	//连接redis数据库
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println("redis err:", err)
	}
	defer conn.Close()

	//操作redis数据库
	reply, err := conn.Do("set", "zxq", "shazi")

	//回复助手类函数   确定成具体的数据类型
	str, err := redis.String(reply, err)

	fmt.Println(str)
}
