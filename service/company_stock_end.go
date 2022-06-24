package service

import (
	"context"
	"log"
	"stock_controller/glob"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CompanyStockEnd struct{}

type CompanyStockEndInterface interface {
}

func NewCompanyStockEnd() *CompanyStockEnd {
	return &CompanyStockEnd{}
}

func (*CompanyStockEnd) CompanyStockEndRun() {
	var stockCompanies []stockCompanyInfo
	findOptions := options.Find()
	stockConnect := glob.MongoDB.Database("stock").Collection("company")
	result, err := stockConnect.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Println(err)
	}

	for result.Next(context.TODO()) {
		var info stockCompanyInfo
		err := result.Decode((&info))
		if err != nil {
			log.Println(err)
		}
		stockCompanies = append(stockCompanies, info)
	}

	result.Close(context.TODO())

}

func (*CompanyStockEnd) EveryDayStockStore(company []stockCompanyInfo) {


	

	urlString := "https://mis.twse.com.tw/stock/api/getStockInfo.jsp?ex_ch=tse_" + stockNumber + ".tw&json=1&delay=0"
}
