/**
* @Author:  Sergey Jay
* @Email :  zshangan@iCloud.com
* @Create:  06/11/2017 15:18
* Copyright (c) 2017 Sergey Jay All rights reserved.
* Description:
 */

package main

import (
	"fmt"
	"net"

	pb "Golang/important/gRPC/example/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	//"google.golang.org/grpc/credentials" // 引入grpc认证包
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata" // 引入grpc meta包
	//"github.com/net/trace"
	"golang.org/x/net/trace"
	"log"
	"net/http"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

// 定义helloService并实现约定的接口
type helloService struct{}

// HelloService ...
var HelloService = helloService{}

var visitTimes int64

func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	visitTimes++
	fmt.Println(visitTimes, "time visit service")

	appid := "101010"
	appkey := "i am key"

	//fmt.Sprintf("Hello %s.\nToken info: appid=%s,appkey=%s", in.Name, appid, appkey)

	resp := new(pb.HelloReply)
	//resp.Message = "Hello " + in.Name + "."	//被覆盖
	resp.Message = fmt.Sprintf("Hello %s.\nToken info: appid=%s,appkey=%s", in.Name, appid, appkey)
	return resp, nil
}

//auth验证token
func auth(ctx context.Context) error {
	// 解析metada中的信息并验证
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return grpc.Errorf(codes.Unauthenticated, "无Token认证信息")
	}

	var (
		appid  string
		appkey string
	)

	if val, ok := md["appid"]; ok {
		appid = val[0]
	}

	if val, ok := md["appkey"]; ok {
		appkey = val[0]
	}

	if appid != "101010" || appkey != "i am key" {
		return grpc.Errorf(codes.Unauthenticated, "Token认证信息无效: appid=%s, appkey=%s", appid, appkey)
	}
	return nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	//
	//var opts []grpc.ServerOption
	//
	//// TLS认证
	//creds, err := credentials.NewServerTLSFromFile("../../keys/server.pem", "../../keys/server.key")
	//if err != nil {
	//	grpclog.Fatalf("Failed to generate credentials %v", err)
	//}
	//opts = append(opts,grpc.Creds(creds))
	//
	//
	//// 注册interceptor
	//var interceptor grpc.UnaryServerInterceptor
	//
	//interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	//	err = auth(ctx)
	//	if err != nil {
	//		return
	//	}
	//	return handler(ctx,req)
	//}
	//opts = append(opts,grpc.UnaryInterceptor(interceptor))

	// 实例化grpc Server, 并开启TLS认证
	s := grpc.NewServer()

	// 注册HelloService
	pb.RegisterHelloServer(s, HelloService)

	go startTrace()

	fmt.Println("Listen on : " + Address)

	s.Serve(listen)
}

func startTrace() {
	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true, true
	}
	go http.ListenAndServe(":50051", nil)
	log.Println("Trace listen on 50051")
}
