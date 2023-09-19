package demo

import (
	"context"
	"helloworld/internal/biz/demo"
)

type DemoImpl struct {
}

func NewDemoImpl() demo.DemoRepo {
	return &DemoImpl{}
}

func (di *DemoImpl) GetDemo(ctx context.Context, conds map[string]interface{}) (*demo.Demo, error) {
	return &demo.Demo{ID: 1}, nil
}
