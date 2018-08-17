package main

import (
	//"google.golang.org/grpc/status"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"os"
	//"net"
	"Golang/important/gRPC/Accounts/rpc"
	"google.golang.org/grpc/credentials"
)

var StoreReq = rpc.StoreReq{
	AccountToken: "inconshreveable",
	Key:          "gopher",
	Val:          []byte("con"),
}

func main() {
	if err := runClient(); err != nil {
		fmt.Fprintf(os.Stderr, "failed: %v\n", err)
		os.Exit(1)
	}
}

func runClient() error {
	tlsCreds, _ := credentials.NewClientTLSFromFile("/Users/shangan/Desktop/GO/src/Golang/important/gRPC/Accounts/keys/Account.crt",
		"myserver")
	conn, err := grpc.Dial("127.0.0.1:8688", grpc.WithTransportCredentials(tlsCreds))
	//  建立连接
	//conn, err := grpc.Dial("localhost:5053", grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to dial server: %v", err)
	}
	defer conn.Close()

	cache := rpc.NewCacheClient(conn)

	//  调用 grpc 的 store() 方法存储键值对 { "gopher": "con" }
	_, err = cache.Store(context.Background(), &StoreReq)
	//_, err = cache.Store(context.Background(), &rpc.StoreReq{
	//	AccountToken: "inconshreveable",
	//	Key:          "gopher",
	//	Val:          []byte("con"),
	//})
	if err != nil {
		//log.Println("Error after cache.Store ~ ")
		return fmt.Errorf("failed to store -> : %v", err)
	}
	//调用 grpc 的 get() 方法取回键为 `gopher` 的值
	resp, err := cache.Get(context.Background(), &rpc.GetReq{Key: "gopher"})
	if err != nil {
		//log.Println("Error after cache.Store !! ")
		return fmt.Errorf("failed to get: %v", err)
	}
	//输出
	fmt.Printf("Got cached value %s\n", resp.Val)
	return nil
}
