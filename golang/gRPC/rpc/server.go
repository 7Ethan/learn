package main

import (
	"log"
	"net"
	"runtime"
	"strconv"

	r "Golang/important/gRPC/rpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	port = "8090"
)

type Data struct{}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// TLS
	tlsCreds, err := credentials.NewServerTLSFromFile("tls.crt", "tls.key")
	if err != nil {
		log.Printf("NewServer Error! %s", err)
	}

	s := grpc.NewServer(grpc.Creds(tlsCreds))

	//起服务
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	r.RegisterDataServer(s, &Data{})
	s.Serve(lis)
	log.Println("grpc server in: %s", port)
}

// 定义方法
func (t *Data) GetUser(ctx context.Context, request *r.UserRq) (response *r.UserRp, err error) {
	response = &r.UserRp{
		Name: strconv.Itoa(int(request.Id)) + ":test",
	}
	return response, err
}
