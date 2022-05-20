package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/afocus/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/consul"
	"iHome/web/model"
	"iHome/web/proto/getCaptcha"
	userMicro "iHome/web/proto/user"
	"iHome/web/utils"
	"image/png"
	"net/http"
)

//获取session信息
func GetSession(ctx *gin.Context) {
	// 初始化错误返回的 map
	resp := make(map[string]interface{})

	//初始化session对象
	s := sessions.Default(ctx)
	userName := s.Get("userName")

	//用户没登录--没存在MySQL中，也没存在session中
	if userName == nil {
		resp["errno"] = utils.RECODE_SESSIONERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
	} else {
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
		var nameData struct {
			Name string `json:"name"`
		}
		nameData.Name = userName.(string) //类型断言
		resp["data"] = nameData
	}

	ctx.JSON(http.StatusOK, resp)
}

//获取图片信息
func GetImageCd(ctx *gin.Context) {
	//获得图片的uuid
	uuid := ctx.Param("uuid")

	//校验数据
	if uuid == "" {
		fmt.Println("获取数据错误")
		return
	}

	//指定服务发现
	//初始化consul
	consulReg := consul.NewRegistry()
	consulSrv := micro.NewService(
		micro.Registry(consulReg),
	)

	//初始化客户端
	microClient := getCaptcha.NewGetCaptchaService("getCaptcha", consulSrv.Client())

	//调用远程函数
	resp, err := microClient.Call(context.TODO(), &getCaptcha.Request{Uuid: uuid})
	if err != nil {
		fmt.Println("未找到远程服务 err:", err)
	}

	//将得到的数据反序列化，得到图片
	var img captcha.Image
	json.Unmarshal(resp.Img, &img)

	//将图片写出到浏览器
	png.Encode(ctx.Writer, img)

	//fmt.Println(str)
	fmt.Println("uuid:", uuid)
}

//发送注册信息
func PostRet(ctx *gin.Context) {

	//ajax传数据需要用以下方法
	//获取数据
	var regData struct {
		Mobile string `json:"mobile"`
		Pwd    string `json:"password"`
	}

	err := ctx.Bind(&regData)
	//校验数据
	if err != nil {
		fmt.Println("获取前段传递数据失败")
		return
	}

	//初始化客户端
	microSrv := utils.InitMicro()
	microClient := userMicro.NewUserService("user", microSrv.Client())

	//调用远程函数
	resp, err := microClient.Register(context.TODO(), &userMicro.Request{
		Mobile:   regData.Mobile,
		Password: regData.Pwd,
	})

	if err != nil {
		fmt.Println("Register err:", err)
		return
	}

	//写给浏览器
	ctx.JSON(http.StatusOK, resp)

	fmt.Println("获取到的数据为：", regData)

	//postform表单传的数据用这种方法获取
	//mobile := ctx.PostForm("mobile")
	//pwd := ctx.PostForm("password")
	//
	//fmt.Println(mobile, pwd)
}

//获得地域信息
func GetArea(ctx *gin.Context) {
	//从缓存redis中取数据
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println("redis err:", err)
		return
	}
	defer conn.Close()

	var areas []model.Area
	//之前使用字节切片存储，现在使用字节切片接受
	areaData, _ := redis.Bytes(conn.Do("get", "areaData"))

	//没有获取到数据
	if len(areaData) == 0 {
		fmt.Println("MySQL获取数据...")
		//先从MySQL中获取数据
		model.GlobalConn.Find(&areas)

		//再把数据写入redis中,存储结构体序列化后的json格式
		areaBuf, _ := json.Marshal(areas)
		conn.Do("set", "areaData", areaBuf)
	} else {
		fmt.Println("Redis获取数据...")
		json.Unmarshal(areaData, &areas)
	}

	resp := make(map[string]interface{})

	resp["errno"] = "0"
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	resp["data"] = areas
	fmt.Println(areas)

	//调用远程服务,获取所有地域信息
	//初始化客户端
	//从consul中获取服务
	//consulRegistry := consul.NewRegistry()
	//micService := micro.NewService(
	//	micro.Registry(consulRegistry),
	//)
	//
	//microClient := getArea.NewGetAreaService("getArea", micService.Client())
	////调用远程服务
	//resp, err := microClient.MicroGetArea(context.TODO(), &getArea.Request{})
	//if err != nil {
	//	fmt.Println("MicroGetArea:", err)
	//	/*ctx.JSON(http.StatusOK,resp)
	//	return */
	//}

	////初始化客户端
	//microSrv := utils.InitMicro()
	//microClient := getArea.NewGetAreaService("go.micro.srv.getArea", microSrv.Client())
	//
	////调用远程服务
	//resp, err := microClient.MicroGetArea(context.TODO(), &getArea.Request{})
	//if err != nil {
	//	fmt.Println("MicroGetArea:", err)
	//	/*ctx.JSON(http.StatusOK,resp)
	//	return */
	//}

	ctx.JSON(http.StatusOK, resp)
}

//处理登录业务
func PostLogin(ctx *gin.Context) {
	//获取前端数据
	var LoginData struct {
		Mobile   string `json:"mobile"`
		PassWord string `json:"password"`
	}
	ctx.Bind(&LoginData)

	resp := make(map[string]interface{})

	//从数据库获取，查询是否匹配
	userName, err := model.Login(LoginData.Mobile, LoginData.PassWord)
	if err != nil {
		fmt.Println("Login err:", err)
		return
	}

	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)

	//将登录状态保存到session中
	//调用session，设置session数据
	s := sessions.Default(ctx)
	//设置session
	s.Set("userName", userName)
	//修改session时，需要Save函数配合，不然不生效
	s.Save()

	ctx.JSON(http.StatusOK, resp)
}

//退出登录
func DeleteSession(ctx *gin.Context) {
	resp := make(map[string]interface{})

	//初始化session对象
	s := sessions.Default(ctx)

	//删除session数据
	s.Delete("userName")

	//必须使用save
	err := s.Save()

	if err != nil {
		fmt.Println("Delete err:", err)
		resp["errno"] = utils.RECODE_IOERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_IOERR)
	} else {
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	}

	ctx.JSON(http.StatusOK, resp)
}

//获取用户信息
func GetUserInfo(ctx *gin.Context) {
	// 初始化错误返回的 map
	resp := make(map[string]interface{})

	//获取session，得到当前用户信息
	//初始化session对象
	s := sessions.Default(ctx)
	userName := s.Get("userName")

	//用户没登录--没存在MySQL中，也没存在session中
	//用户没登录，但是进入该页面，恶意进入
	if userName == nil {
		resp["errno"] = utils.RECODE_SESSIONERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
	} else {
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)

		//根据用户名，从MySQL中查找用户详细信息
		user, err := model.GetUserInfo(userName.(string))
		if err != nil {
			fmt.Println("GetUserInfo MySQL err:", err)
			return
		}

		temp := make(map[string]interface{})
		temp["user_id"] = user.ID
		temp["name"] = user.Name
		temp["mobile"] = user.Mobile
		temp["real_name"] = user.Real_name
		temp["id_card"] = user.Id_card
		temp["avatar_url"] = user.Avatar_url

		resp["data"] = temp
	}

	ctx.JSON(http.StatusOK, resp)
}

//更新用户名
func PutUserInfo(ctx *gin.Context) {
	resp := make(map[string]interface{})
	defer ctx.JSON(http.StatusOK, resp)

	//获取当前用户名
	s := sessions.Default(ctx)
	userName := s.Get("userName")

	//获取新用户名-----处理Request Rayload类型数据，Bind()
	var nameData struct {
		Name string `json:"name"`
	}
	ctx.Bind(&nameData)

	//更新用户名
	//更新数据库中的name
	err := model.UpdateUserName(userName.(string), nameData.Name)

	if err != nil {
		fmt.Println("UpdateUserName err:", err)
		resp["errno"] = utils.RECODE_DBERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
		return
	}

	//更新session数据
	s.Set("userName", nameData.Name)
	err = s.Save()
	if err != nil {
		fmt.Println("session err:", err)
		resp["errno"] = utils.RECODE_SESSIONERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
		return
	}
	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	resp["data"] = nameData
}

//上传头像
func PostAvatar(ctx *gin.Context) {
	//获取图片文件，静态文件对象
	file, _ := ctx.FormFile("avatar")

	//存储上传文件到项目中
	err := ctx.SaveUploadedFile(file, "web/view/head portrait/"+file.Filename)
	fmt.Println("SaveUploadedFile err:", err)

	resp := make(map[string]interface{})
	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)

	//temp := make(map[string]interface{})
	resp["data"] = file

	ctx.JSON(http.StatusOK, resp)
}

type AuthStu struct {
	IdCard   string `json:"id_card"`
	RealName string `json:"real_name"`
}

//上传实名认证
func PostUserAuth(ctx *gin.Context) {
	//获取数据
	var auth AuthStu
	err := ctx.Bind(&auth)

	if err != nil {
		fmt.Println("AuthStu err:", err)
		return
	}

	//session := sessions.Default(ctx)
	//userName := session.Get("userName")

	//处理数据 微服务
	//microClient := user.NewUserService("postUserAuth", utils.Get)

	//调用远程服务

	//返回数据

}
