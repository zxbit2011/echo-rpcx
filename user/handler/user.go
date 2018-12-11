package handler

import "context"

type User int

type Account struct {
	Mobile   string
	Name     string
	Sex      string
	Nickname string
}

type Result struct {
	Code int
	Msg  string
	Data interface{}
}

func (u User) Info(ctx context.Context, token string, res *Result) (err error) {
	if token == "token" {
		res.Code = 200
		res.Msg = "获取成功"
		res.Data = &Account{
			Mobile:   "18223200000",
			Name:     "zxbit2011",
			Sex:      "男",
			Nickname: "承诺一时的华丽",
		}
		return
	}
	res.Code = 401
	res.Msg = "获取用户信息失败"
	return
}
