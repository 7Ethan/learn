package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	//_ "github.com/grpc/grpc-go/codes"
	"Golang/important/gRPC/Accounts/rpc"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"os"
	"runtime"
)

type CacheService struct {
	accounts      rpc.AccountsClient
	store         map[string][]byte
	keysByAccount map[string]int64
}

var storeTimes int64

func (s *CacheService) Get(ctx context.Context, req *rpc.GetReq) (*rpc.GetResp, error) {
	val, ok := s.store[req.Key]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "Key not found %s", req.Key)
	}
	return &rpc.GetResp{Val: val}, nil
}

//func (s *CacheService)GetByToken(ctx context.Context,req *rpc.GetByTokenReq)(*rpc.GetByTokenResp,error)  {
//	//resp := new(rpc.GetByTokenResp)
//	Account := &rpc.Account{100}
//	return &rpc.GetByTokenResp{Account},nil
//}

func (s *CacheService) Store(ctx context.Context, req *rpc.StoreReq) (*rpc.StoreResp, error) {
	storeTimes++
	fmt.Println(storeTimes, "time to visit store")

	//=========
	s.store = make(map[string][]byte)
	s.keysByAccount = make(map[string]int64)
	//s.accounts = rpc.Account{MaxCacheKeys:s.keysByAccount[req.Key]}
	//  调用另一个服务取得账户信息，包含其键值限制
	resp, err := s.accounts.GetByToken(context.Background(), &rpc.GetByTokenReq{
		Token: req.AccountToken,
	})
	//resp, err := s.accounts.GetByToken(context.Background(), &GetByTokenReq)

	if err != nil {
		return nil, err
	}
	//// 检查是否超量使用
	if s.keysByAccount[req.AccountToken] >= resp.Account.MaxCacheKeys {
		return nil, status.Errorf(codes.FailedPrecondition, "Account %s exceeds max key limit %d", req.AccountToken, resp.Account.MaxCacheKeys)
	}
	//  如果键不存在，需要新加键值，那么我们就对计数器加一
	if _, ok := s.store[req.Key]; !ok {
		s.keysByAccount[req.AccountToken] += 1
	}

	//  保存键值
	s.store[req.Key] = req.Val
	return &rpc.StoreResp{}, nil
}

//var Service = new(CacheService)

var Service = CacheService{}

func runServer() error {
	tlsCreds, err := credentials.NewServerTLSFromFile("/Users/shangan/Desktop/GO/src/Golang/important/gRPC/Accounts/keys/Account.crt",
		"/Users/shangan/Desktop/GO/src/Golang/important/gRPC/Accounts/keys/Account.key")
	if err != nil {
		log.Printf("Newserver err : %s \n", err)
		return err
	}
	srv := grpc.NewServer(grpc.Creds(tlsCreds))

	rpc.RegisterCacheServer(srv, &Service)
	l, err := net.Listen("tcp", "127.0.0.1:8688")
	if err != nil {
		log.Printf("listen err : %s \n", err)
		return err
	}
	//  block
	return srv.Serve(l)
}

func main() {
	runtime.GOMAXPROCS(4)
	if err := runServer(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run cache server: %s\n", err)
		//os.Exit(1)
	}
}

/*
定义服务(Service)
如果想要将消息类型用在RPC(远程方法调用)系统中，可以在.proto文件中定义一个RPC服务接口，protocol buffer编译器会根据所选择的不同语言生成服务接口代码。例如，想要定义一个RPC服务并具有一个方法，该方法接收SearchRequest并返回一个SearchResponse，此时可以在.proto文件中进行如下定义：

service SearchService {
    rpc Search (SearchRequest) returns (SearchResponse) {}
}
生成的接口代码作为客户端与服务端的约定，服务端必须实现定义的所有接口方法，客户端直接调用同名方法向服务端发起请求。比较蛋疼的是即便业务上不需要参数也必须指定一个请求消息，一般会定义一个空message。


*/
