package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/labstack/echo"
	"github.com/smallnest/rpcx/client"
	"github.com/zxbit2011/echo-rpcx/sso/handler"
	"log"
	"net/http"
)

var (
	ssoAddr  = flag.String("ssoAddr", "127.0.0.1:5001", "server address")
	userAddr = flag.String("userAddr", "127.0.0.1:5002", "server address")
)

func main() {
	//单点登录服务
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}
	option := client.DefaultOption
	option.TLSConfig = conf
	d := client.NewPeer2PeerDiscovery("quic@"+*ssoAddr, "")
	ssoClient := client.NewXClient("SSO", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer ssoClient.Close()

	//用户中心

	d = client.NewPeer2PeerDiscovery("tcp@"+*userAddr, "")
	userClient := client.NewXClient("User", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer userClient.Close()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World - Echo Web!")
	})
	e.GET("/login", func(c echo.Context) error {
		login := &handler.Login{
			Mobile:   c.FormValue("mobile"),
			Password: c.FormValue("password"),
			Code:     c.FormValue("code"),
		}
		res := &handler.Result{}
		err := ssoClient.Call(context.Background(), "Login", login, res)
		if err != nil {
			log.Fatalf("【sso login】failed to call: %v", err)
			return c.String(http.StatusOK, fmt.Sprintf("【sso login】failed to call: %v", err))
		}

		log.Printf("Result：Code:%v，Msg：%s，Data，%v", res.Code, res.Msg, res.Data)
		return c.JSON(http.StatusOK, res)
	})
	e.GET("/user/info", func(c echo.Context) error {
		res := &handler.Result{}
		err := userClient.Call(context.Background(), "Info", "token", res)
		if err != nil {
			log.Fatalf("【user info】failed to call: %v", err)
			return c.String(http.StatusOK, fmt.Sprintf("【user info】failed to call: %v", err))
		}

		log.Printf("Result：Code:%v，Msg：%s，Data，%v", res.Code, res.Msg, res.Data)
		return c.JSON(http.StatusOK, res)
	})
	e.Logger.Fatal(e.Start(":5000"))
}
