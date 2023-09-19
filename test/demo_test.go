package test

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	v1 "helloworld/api/helloworld/v1"
	"testing"
)

func getRpcClient() (*grpc.ClientConn, error) {
	rpcClient, err := grpc.Dial("127.0.0.1:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(fmt.Sprintf("get rpcClient error = %v", err))
		return nil, err
	}
	return rpcClient, nil
}

func TestGetDemo(t *testing.T) {
	client, err := getRpcClient()
	if err != nil {
		fmt.Println(fmt.Sprintf("get rpcClient error = %v", err))
		return
	}
	data, err := v1.NewDemoClient(client).SayDemo(context.Background(), &v1.Trades{Id: int32(1)})
	if err != nil {
		fmt.Println(fmt.Sprintf("get data error = %v", err))
		return
	}
	fmt.Println(data)
}
