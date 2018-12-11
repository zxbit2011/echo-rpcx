package main

import (
	"crypto/tls"
	"flag"
	"github.com/smallnest/rpcx/server"
	"github.com/zxbit2011/echo-rpcx/sso/handler"
	"log"
)

var (
	addr = flag.String("addr", "localhost:5001", "server address")
)

func main() {
	flag.Parse()
	cert, err := tls.LoadX509KeyPair("server.pem", "server.key")
	if err != nil {
		log.Print(err)
		return
	}
	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	s := server.NewServer(server.WithTLSConfig(config))
	s.RegisterName("SSO", new(handler.SSO), "")
	err = s.Serve("quic", *addr)
	if err != nil {
		panic(err)
	}
}
