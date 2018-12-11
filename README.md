# echo + rpcx
```
这是基于echo、rpcx框架的demo ，基于rpcx展示了quic,tcp协议使用方式。
```
` 注意条件编译，否则quic协议无法运行 `
##### sso 登录服务（` $ cd sso `）：
```
$ go run -tags "quic" main.go
```
##### user 用户中心服务（` $ cd user `）：
```
$ go run main.go
```
##### web（` $ cd web `）：
```
$ go run -tags "quic" main.go
```
