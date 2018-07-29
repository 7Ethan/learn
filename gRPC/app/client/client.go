package main

import (
	//"google.golang.org/grpc/status"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	//_ "github.com/grpc/grpc-go/codes"
	"fmt"
	//"net"
	"Golang/important/gRPC/app/rpc"
	//"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials"
	"log"
)

//  调用 grpc 的 store() 方法存储键值对 { "gopher": "con" }
var StoreReq = rpc.StoreReq{Key: "gopher", Val: "SHIIIIIIIITTTT!!!!"}

func main() {
	var err error
	var opts []grpc.DialOption

	creds, err := credentials.NewClientTLSFromFile("/Users/shangan/Desktop/GO/src/Golang/important/gRPC/app/keys/server.pem",
		"MyServer")
	if err != nil {
		fmt.Printf("Failed to create TLS credentials %v", err)
	}

	opts = append(opts, grpc.WithTransportCredentials(creds))

	conn, err := grpc.Dial("localhost:8088", opts...)
	if err != nil {
		fmt.Printf("Failed to dial %v", err)
	}
	defer conn.Close()

	cache := rpc.NewCacheClient(conn)

	_, err = cache.Store(context.Background(), &StoreReq)
	if err != nil {
		log.Println("error after cache.Store ~ ")
		//return fmt.Errorf("failed to store: %v", err)
	}
	//  调用 grpc 的 get() 方法取回键为 `gopher` 的值
	resp, err := cache.Get(context.TODO(), &rpc.GetReq{Key: "gopher"})
	if err != nil {
		log.Printf("error after cache.Get ~ ERROR : %v", err)
		//return fmt.Errorf("failed to get: %v", err)
	}
	//  输出
	fmt.Printf("Got cached value %v\n", resp)

}

//
//func runClient() error {
//	//tlsCreds, _ := credentials.NewClientTLSFromFile("/Users/shangan/Desktop/GO/src/Golang/important/gRPC/app/keys/server.pem",
//	//	"MyServer")
//	//conn, err := grpc.Dial("localhost:5051", grpc.WithTransportCredentials(tlsCreds))
//	//  建立连接
//	conn, err := grpc.Dial("localhost:5051", grpc.WithInsecure())
//	if err != nil {
//		return fmt.Errorf("failed to dial server: %v", err)
//	}
//	defer conn.Close()
//
//	cache := rpc.NewCacheClient(conn)
//
//	//  调用 grpc 的 store() 方法存储键值对 { "gopher": "con" }
//	storeReq := rpc.StoreReq{Key: "gopher", Val:[]byte("con")}
//	_, err = cache.Store(context.Background(), &storeReq)
//	if err != nil {
//		log.Println("err after cache.Store ~ ")
//		return fmt.Errorf("failed to store: %v", err)
//	}
//	//  调用 grpc 的 get() 方法取回键为 `gopher` 的值
//	resp, err := cache.Get(context.Background(), &rpc.GetReq{Key: "gopher"})
//	if err != nil {
//		return fmt.Errorf("failed to get: %v", err)
//	}
//	//  输出
//	fmt.Printf("Got cached value %s\n", resp.Val)
//	return nil
//
//}
