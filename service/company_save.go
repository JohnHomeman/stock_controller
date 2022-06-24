package service

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"stock_controller/glob"

	"github.com/henrylee2cn/surfer"
	"github.com/robfig/cron/v3"
)

const (
	checkDay = "0 14 * * *"
)

type CompanySave struct{}

type CompanySaveInterface interface {
	companySaveRun()
}

type stockCompanyInfo struct {
	CompanyAllName  string `json:"companyAllName"`
	CompanyNickName string `json:"companyNickName"`
	CompanyNumber   string `json:"companyNumber"`
	CompanyCatagory string `json:"companyCatagory"`
}

func NewCompanySave() *CompanySave {
	return &CompanySave{}
}

func (*CompanySave) CompanySaveRun() {

	searchStockData()

	c := cron.New()
	c.AddFunc(checkDay, func() {
		searchStockData()
	})
	c.Start()

}

func searchStockData() {
	urlString := "https://openapi.twse.com.tw/v1/opendata/t187ap03_L"
	req := &surfer.Request{
		Method: "Get",
		Url:    urlString,
	}
	resp, err := surfer.Download(req)
	if err != nil {
		log.Println("API URL Error, " + err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Body Error")

	}
	var sourceData []interface{}
	if err := json.Unmarshal(body, &sourceData); err != nil {
		log.Println("Unmarshal Error")
	}
	for _, p := range sourceData {
		temp := p.(map[string]interface{})
		insertInform := stockCompanyReturn(temp["公司名稱"].(string), temp["公司簡稱"].(string), temp["公司代號"].(string), temp["產業別"].(string))
		stockConnect := glob.MongoDB.Database("stock").Collection("company")
		result, err := stockConnect.InsertOne(context.TODO(), insertInform)
		if err != nil {
			panic(err)
		}
		log.Println(result.InsertedID)

	}
}

func stockCompanyReturn(companyAllName string, companyNickName string, companyNumber string, companyCatagory string) *stockCompanyInfo {
	return &stockCompanyInfo{
		CompanyAllName:  companyAllName,
		CompanyNickName: companyNickName,
		CompanyNumber:   companyNumber,
		CompanyCatagory: companyCatagory,
	}
}
