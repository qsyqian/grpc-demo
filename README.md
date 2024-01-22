### protoc

```shell
# protoc --version
libprotoc 3.21.12
```

安装参考：https://grpc.io/docs/protoc-installation/

### go-grpc相关工具

```shell
# ll $GOPATH/bin  
total 178752
-rwxr-xr-x  1 x  staff    59M Jan 18 17:54 goctl
-rwxr-xr-x  1 x  staff   8.7M Jan 18 10:13 protoc-gen-go
-rwxr-xr-x  1 x  staff   8.5M Jan 18 10:15 protoc-gen-go-grpc
-rwxr-xr-x  1 x  staff    11M Jan 19 15:45 protoc-gen-grpc-gateway
```

安装参考：https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/introduction/
```shell
$ go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### genprotoc.sh

这个命令的位置要对，比如本例子中，该命令在pb下，所以他后面跟的参数都是从当前目录开始算起。

执行的时候也许要进入到pb目录下执行。

### run

启动server，进入到server目录执行

```shell
# go run server.go
2024/01/22 15:29:57 Serving on http://127.0.0.1:8091
```
执行请求，进入到client目录执行

```shell
# curl -X POST -k http://127.0.0.1:8091/v1/sayHello -d '{"name": "Curl Client"}'
{"message":"Hi: Curl Client"}%
# go run client.go
2024/01/22 15:31:08 get msg:  Hi: Grpc Client
```

server日志
```shell
# go run server.go
2024/01/22 15:29:57 Serving on http://127.0.0.1:8091
2024/01/22 15:31:01 get msg: name:"Curl Client"
2024/01/22 15:31:08 get msg: name:"Grpc Client"
```