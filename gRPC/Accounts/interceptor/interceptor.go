/**
* @Author:  Sergey Jay
* @Email :  zshangan@iCloud.com
* @Create:  09/11/2017 09:48
* Copyright (c) 2017 Sergey Jay All rights reserved.
* Description:
 */

package interceptor

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func WithClientInterceptor() grpc.DialOption {
	return grpc.WithUnaryInterceptor(clientInterceptor)
}
func clientInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	log.Printf("invoke remote method=%s duration=%s error=%v", method, time.Since(start), err)
	return err
}
