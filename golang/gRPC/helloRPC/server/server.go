package main

// server.go

import (
	"net"

	pb "Golang/important/gRPC/helloRPC/helloworld"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

const (
	Address = "127.0.0.1:50052"
)

var visitTimes int64

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	visitTimes++
	fmt.Printf("It's %d time's visit~\n", visitTimes)
	return &pb.HelloReply{Message: "Hello " + in.Name + "."}, nil
}

func main() {
	lis, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	//TLS认证
	creds, err := credentials.NewServerTLSFromFile("/Users/shangan/Desktop/GO/src/Golang/important/gRPC/helloRPC/keys/server.pem",
		"/Users/shangan/Desktop/GO/src/Golang/important/gRPC/helloRPC/keys/server.key")
	if err != nil {
		grpclog.Fatalf("Failed to generate credentials %v", err)
	}

	// 实例化grpc Server, 并开启TLS认证
	s := grpc.NewServer(grpc.Creds(creds))

	// 注册HelloService
	pb.RegisterGreeterServer(s, &server{})
	grpclog.Println("Listen on " + Address)
	s.Serve(lis)
}
