package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	//_ "github.com/grpc/grpc-go/codes"
	"Golang/important/gRPC/app/rpc"
	"fmt"
	"google.golang.org/grpc/codes"
	"net"
	//"log"
	//"github.com/coreos/etcd/pkg/srv"
	"google.golang.org/grpc/credentials"
)

type CacheService struct {
	store map[string]string
}

var Service = CacheService{}

var storeTimes int64

func (s *CacheService) Get(ctx context.Context, req *rpc.GetReq) (*rpc.GetResp, error) {
	val, ok := s.store[req.Key]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "Key not found %s", req.Key)
	}
	return &rpc.GetResp{Val: val}, nil
}

func (s *CacheService) Store(ctx context.Context, req *rpc.StoreReq) (*rpc.StoreResp, error) {
	s.store = make(map[string]string)
	s.store[req.Key] = req.Val

	storeTimes++
	fmt.Println(storeTimes, "times request store #  ")
	return &rpc.StoreResp{}, nil
}

func main() {
	listen, err := net.Listen("tcp", "localhost:8088")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile("/Users/shangan/Desktop/GO/src/Golang/important/gRPC/app/keys/server.pem",
		"/Users/shangan/Desktop/GO/src/Golang/important/gRPC/app/keys/server.key")
	if err != nil {
		fmt.Printf("Failed to generate credentials %v", err)
	}

	s := grpc.NewServer(grpc.Creds(creds))

	rpc.RegisterCacheServer(s, &Service)

	s.Serve(listen)
}

//
//func runServer() error {
//	//tlsCreds, err := credentials.NewServerTLSFromFile("/Users/shangan/Desktop/GO/src/Golang/important/gRPC/app/keys/server.pem",
//	//	"/Users/shangan/Desktop/GO/src/Golang/important/gRPC/app/keys/server.key")
//	//if err != nil {
//	//	log.Printf("Newserver err : %s \n", err)
//	//	return err
//	//}
//	//srv := grpc.NewServer(grpc.Creds(tlsCreds))
//
//	srv := grpc.NewServer()
//	//service := new(CacheService)
//
//	rpc.RegisterCacheServer(srv, &Service)
//	l, err := net.Listen("tcp", "localhost:5051")
//	if err != nil {
//		log.Printf("listen err : %s \n", err)
//		return err
//	}
//	//  block
//	return srv.Serve(l)
//}
