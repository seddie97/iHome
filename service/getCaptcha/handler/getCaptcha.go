package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"getCaptcha/model"
	"github.com/afocus/captcha"
	"image/color"

	getCaptcha "getCaptcha/proto"
)

type GetCaptcha struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *GetCaptcha) Call(ctx context.Context, req *getCaptcha.Request, rsp *getCaptcha.Response) error {

	//生成验证码
	// 初始化对象
	cap := captcha.New()

	// 设置字体
	cap.SetFont("./config/comic.ttf")

	// 设置验证码大小
	cap.SetSize(128, 64)

	// 设置干扰强度
	cap.SetDisturbance(captcha.HIGH)

	// 设置前景色
	cap.SetFrontColor(color.RGBA{R: 255, G: 245, B: 247, A: 100})

	// 设置背景色
	cap.SetBkgColor(color.RGBA{R: 178, G: 200, B: 187, A: 100}, color.RGBA{R: 214, G: 200, B: 75, A: 100}, color.RGBA{R: 167, G: 220, B: 224, A: 100})

	// 生成字体 -- 将图片验证码, 展示到页面中.
	img, str := cap.Create(4, captcha.ALL)

	//将验证码存入redis数据库
	fmt.Println(str)
	err := model.SaveImgCode(req.Uuid, str)
	if err != nil {
		fmt.Println("microClient err:", err)
		return err
	}

	//将生成的图片序列化
	imgBuf, _ := json.Marshal(img)

	//将imgBuf 使用参数resp传出
	rsp.Img = imgBuf

	return nil
}
