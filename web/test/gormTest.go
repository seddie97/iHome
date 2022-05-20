package main

import (
	"fmt"
	"github.com/micro/go-micro/sync/time"

	//"_"导入驱动包的时候下划线表示：
	//有init函数，注册mysql驱动
	//orm框架连接mysql驱动，在项目入口main函数启动之前被调用
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	//_ "iHome/web/test2"
)

type Student struct {
	//gorm.Model //go语言中，匿名结构体成员-----继承！
	Id int
	//string -- varchar  默认大小255，可以在创建表时，指定大小
	Name  string `gorm:"size:100;default:'xiaoming'"`
	Age   int
	Class int       `gorm:"not null"`
	Join  time.Time `gorm:"type:datetime"`
}

// GlobalConn 创建全局的连接池句柄
var GlobalConn *gorm.DB

func main() {
	//fmt.Println("这是main函数")
	//连接数据库--获取连接池的句柄	格式：用户名:密码@协议(IP:port)/数据库名
	conn, err := gorm.Open("mysql",
		"root:123456@tcp(127.0.0.1:3306)/test?parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("gorm conn err:", err)
		return
	}
	//defer conn.Close()
	GlobalConn = conn

	//初始数
	GlobalConn.DB().SetMaxIdleConns(10)
	//最大数
	GlobalConn.DB().SetMaxOpenConns(100)

	//创建非复数表名的数据库表
	GlobalConn.SingularTable(true)
	//借助gorm创建数据库表
	fmt.Println(GlobalConn.AutoMigrate(new(Student)).Error)

	//插入数据
	//InsertData()

	//查询数据
	//SearchData()

	//更新数据
	//UpdateData()

	//删除数据
	//DeleteData()
}

func InsertData() {
	//先创建数据---创建结构体对象
	var stu Student
	stu.Name = "ldj"
	stu.Age = 24

	//插入（add）数据
	fmt.Println(GlobalConn.Create(&stu).Error)
}

func SearchData() {
	var stu []Student
	//查询第一条的全部信息
	//GlobalConn.First(&stu)

	//只查询第一条的name和age属性
	//GlobalConn.Select("name, age").First(&stu)

	//查询全部的
	//GlobalConn.Find(&stu)

	//条件查询，姓名为ldj的姓名年龄
	//GlobalConn.Select("name, age").Where("name = ?", "zhangdan").Find(&stu)

	//多个条件查询
	//GlobalConn.Select("name, age").Where("name = ?", "ldj").Where("age = ?", "24").Find(&stu)

	//多个条件查询合并条件
	//GlobalConn.Select("name, age").Where("name = ? and age = ?", "ldj", 24).Find(&stu)
	GlobalConn.Unscoped().Find(&stu)
	fmt.Println(stu)
}

func UpdateData() {
	var stu Student
	stu.Name = "123"
	stu.Age = 123

	//fmt.Println(GlobalConn.Model(new(Student)).Where("name = ?", "ldj").
	//	Update("name", "lindejia").Error)

	fmt.Println(GlobalConn.Model(new(Student)).Where("name = ?", "lindejia").
		Update(map[string]interface{}{"name": "ldj", "age": 23}).Error)

}

func DeleteData() {
	fmt.Println(GlobalConn.Unscoped().Where("name = ?", "zhangdan").Delete(new(Student)).Error)
}
