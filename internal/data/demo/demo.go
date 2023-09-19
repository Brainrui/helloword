package demo

import (
	"context"
	"helloworld/internal/biz/demo"
)

type DemoRepo struct {
}

func NewDemoRepo() demo.DemoRepo {
	return &DemoRepo{}
}

func (di *DemoRepo) GetDemo(ctx context.Context, conds map[string]interface{}) (*demo.Demo, error) {
	return &demo.Demo{ID: 1}, nil
}
