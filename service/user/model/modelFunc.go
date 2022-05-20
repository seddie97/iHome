package model

import (
	"crypto/md5"
	"encoding/hex"
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

//注册用户信息
func RegisterUser(mobile, pwd string) error {
	var user User
	user.Name = mobile //默认手机作为用户名
	user.Mobile = mobile

	//需要把密码加密
	m5 := md5.New()                       //初始化md5对象
	m5.Write([]byte(pwd))                 //pwd写入缓冲区
	pwd = hex.EncodeToString(m5.Sum(nil)) //不使用额外的秘钥

	user.Password_hash = pwd

	//插入数据库
	return GlobalConn.Create(&user).Error
}
