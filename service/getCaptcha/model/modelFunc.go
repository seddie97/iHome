package model

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

//存储图片id到redis数据库中
func SaveImgCode(uuid, code string) error {
	//连接redis数据库
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println("redis err:", err)
		return err
	}
	defer conn.Close()

	//操作redis数据库--写		有效时间	5	分钟
	_, err = conn.Do("setex", uuid, 60*5, code)

	return err //不需要回复助手！
}
