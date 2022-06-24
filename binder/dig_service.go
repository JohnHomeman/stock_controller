package binder

import (
	"stock_controller/service"
	"sync"

	"stock_controller/glob"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/dig"
)

var (
	backGroundAPP *BackGroundAPP
	serverSetOnce sync.Once
)

type BackGroundAPP struct {
	dig.In
	MongoDB        *mongo.Client
	CompanyService *service.CompanySave
	Plo            *service.FetcherPlo
	Fpln           *service.FetcherFpln
}

func InitBackGroundAPP(app BackGroundAPP) {
	serverSetOnce.Do(func() {
		backGroundAPP = &app
	})

}

func Run() {
	var err error
	binder := New()
	if err = binder.Invoke(InitBackGroundAPP); err != nil {
		panic(err)
	}
	glob.MongoDB = backGroundAPP.MongoDB
	go backGroundAPP.CompanyService.CompanySaveRun()
	service.PloController = backGroundAPP.Plo
	service.FplnController = backGroundAPP.Fpln

}
