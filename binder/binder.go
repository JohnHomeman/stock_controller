package binder

import (
	"stock_controller/glob"
	"stock_controller/service"
	"sync"

	"go.uber.org/dig"
)

var (
	binder *dig.Container
	once   sync.Once
)

func New() *dig.Container {
	once.Do(func() {
		binder = dig.New()
		binder.Provide(glob.NewMongoCfg)
		binder.Provide(glob.NewMongoServer)
		binder.Provide(service.NewCompanySave)
	})
	return binder
}
