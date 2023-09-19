package demo

import (
	"context"
	"fmt"
)

type Demo struct {
	ID int `json:"id" gorm:"id"`
}

func (Demo) TableName() string {
	return ""
}

type DemoRepo interface {
	GetDemo(ctx context.Context, conds map[string]interface{}) (*Demo, error)
}

type DemoUseCase struct {
	di DemoRepo
}

func NewDemoUseCase(di DemoRepo) *DemoUseCase {
	return &DemoUseCase{
		di: di,
	}
}

func (dc *DemoUseCase) GetTrades(ctx context.Context) (*Demo, error) {
	data, err := dc.di.GetDemo(ctx, nil)
	if err != nil {
		fmt.Println("this is a error")
		return nil, err
	}
	return data, nil
}
