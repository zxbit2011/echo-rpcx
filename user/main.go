package main

import (
	"flag"
	"github.com/smallnest/rpcx/server"
	"github.com/zxbit2011/echo-rpcx/user/handler"
)

var (
	addr = flag.String("addr", "localhost:5002", "server address")
)

func main() {
	flag.Parse()
	s := server.Server{}
	s.RegisterName("User", new(handler.User), "")
	go s.Serve("tcp", *addr)
	select {}
}
