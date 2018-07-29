package main

//client.go

import (
	"log"
	"os"

	pb "Golang/important/gRPC/helloRPC/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

const (
	address     = "127.0.0.1:50052"
	defaultName = "world"
)

func main() {
	// TLS连接
	creds, err := credentials.NewClientTLSFromFile("/Users/shangan/Desktop/GO/src/Golang/important/gRPC/helloRPC/keys/server.pem",
		"MyServer")
	if err != nil {
		grpclog.Fatalf("Failed to create TLS credentials %v", err)
	}

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
