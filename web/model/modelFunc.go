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

//处理登录业务,根据手机号和密码获取当前用户名
func Login(mobile, pwd string) (string, error) {
	var user User
	//对参数的pwd作md5的哈希处理
	m5 := md5.New()
	m5.Write([]byte(pwd))
	pwdHash := hex.EncodeToString(m5.Sum(nil))
	err := GlobalConn.Where("mobile = ?", mobile).
		Where("password_hash = ?", pwdHash).Select("name").
		Find(&user).Error

	return user.Name, err
}

//获取用户信息
func GetUserInfo(userName string) (User, error) {
	//实现SQL：select * from user where name = #userName;
	var user User

	err := GlobalConn.Where("name = ?", userName).First(&user).Error

	return user, err
}

//更新用户名
func UpdateUserName(oldName, newName string) error {
	//update user set name ='ldj' where name = '18106939917'
	return GlobalConn.Model(new(User)).Where("name = ?", oldName).Update("name", newName).Error
}
