package handler

import "context"

type SSO int

type Login struct {
	Mobile   string
	Password string
	Code     string
}

type Result struct {
	Code int
	Msg  string
	Data interface{}
}

func (sso SSO) Login(ctx context.Context, login *Login, res *Result) (err error) {
	if login.Mobile == "18223200000" && login.Password == "123456" && login.Code == "1234" {
		res.Code = 200
		res.Msg = "登录成功"
		res.Data = "token"
		return
	}
	res.Code = 401
	res.Msg = "登录失败，手机号或密码错误"
	return
}
