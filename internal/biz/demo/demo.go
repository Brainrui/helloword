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

type DemoCase struct {
	di DemoRepo
}

func NewDemoCase(di DemoRepo) *DemoCase {
	return &DemoCase{
		di: di,
	}
}

func (dc *DemoCase) GetTrades(ctx context.Context) (*Demo, error) {
	data, err := dc.di.GetDemo(ctx, nil)
	if err != nil {
		fmt.Println("this is a error")
		return nil, err
	}
	return data, nil
}
