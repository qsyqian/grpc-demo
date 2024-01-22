package main

import (
	"context"
	"github.com/qsyqian/grpc-demo/pb/hello"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8091", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("grpc dial err: ", err)
	}
	helloClient := hello.NewGreeterClient(conn)
	resp, err := helloClient.SayHello(context.Background(), &hello.HelloRequest{Name: "Grpc Client"})
	if err != nil {
		log.Fatalln("grpc say hello err: ", err)
	}
	log.Println("get msg: ", resp.Message)
}
