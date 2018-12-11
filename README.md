# echo + rpcx
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
