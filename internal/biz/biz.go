package biz

import (
	"github.com/google/wire"
	"helloworld/internal/biz/demo"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewGreeterUsecase, demo.NewDemoUseCase)
