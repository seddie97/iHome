package handler

import (
	"context"
	"user/model"
	user "user/proto"
	"user/utils"
)

type User struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *User) Register(ctx context.Context, req *user.Request, rsp *user.Response) error {
	//先校验短信验证码，redis中的
	//注册用户，将数据写入Mysql数据库
	err := model.RegisterUser(req.Mobile, req.Password)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
	} else {
		rsp.Errno = utils.RECODE_OK
		rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	}

	return nil
}
