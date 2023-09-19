package service

import (
	"context"
	"fmt"
	v1 "helloworld/api/helloworld/v1"
	"helloworld/internal/biz"
	"helloworld/internal/biz/demo"
)

type DemoService struct {
	v1.UnimplementedGreeterServer

	uc  *biz.GreeterUsecase
	duc *demo.DemoCase
}

func NewDemoService(uc *biz.GreeterUsecase, duc *demo.DemoCase) *DemoService {
	return &DemoService{
		uc:  uc,
		duc: duc,
	}
}
func (d *DemoService) SayDemo(ctx context.Context, req *v1.Trades) (*v1.Trades, error) {
	data, err := d.duc.GetTrades(ctx)
	if err != nil {
		fmt.Println(fmt.Sprintf("this is an error = %v", err))
		return nil, err
	}
	res := &v1.Trades{
		Id:   int32(data.ID),
		Name: "name",
	}
	return res, nil
}
