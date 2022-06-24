package model

import (
	"time"
)

type Proxy struct {
	//mongo.CollectionBase `bson:",inline"`
	IPAddress        string    `json:"ip_address" bson:"ip_address"`
	Port             string    `json:"port" bson:"port" `
	SourceCode       string    `json:"source_code" bson:"source_code"`
	Enable           int       `json:"enable" bson:"enable"`
	PassSites        string    `json:"pass_sites" bson:"pass_sites"`
	IsHealthy        bool      `json:"is_healthy" bson:"is_healthy"`
	HTTPS            bool      `json:"https" bson:"https"`
	IsHealthyArchive bool      `json:"is_healthy_archive" bson:"is_healthy_archive"`
	CountryCode      string    `json:"country_code" bson:"country_code"`
	Country          string    `json:"country" bson:"country"`
	Anonymous        bool      `json:"anonymous" bson:"anonymous"`
	UsedCount        int       `json:"used_count" bson:"used_count"`
	CheckTime        time.Time `json:"check_time" bson:"check_time"`
	LastSuccessTime  time.Time `json:"last_success_time" bson:"last_success_time"`
	AliveMin         int       `json:"alive_min" bson:"alive_min"`
	// 上次成功的持續時間 = 這次檢查時間CheckTime-上次成功時間LastSuccessTime
	LastSuccessIntervalHour int       `json:"last_success_interval_hour" bson:"last_success_interval_hour"`
	SuccessCountPerDay      int       `json:"success_count_per_day" bson:"success_count_per_day"`
	CheckRoundPerDay        int       `json:"check_round_per_day" bson:"check_round_per_day"`
	CreateTime              time.Time `json:"create_time" bson:"create_time"`
	Description             string    `json:"description" bson:"description"`
}
