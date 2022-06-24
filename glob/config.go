package glob

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var (
	YamlC = YamlConfig{}
)

type YamlConfig struct {
	MongoDBCfg     MongoDBCfg     `mapstructure:"mongodb"`
	FetcherploCfg  FetcherploCfg  `mapstructure:"fetcherplo"`
	FetcherfplnCfg FetcherfplnCfg `mapstructure:"fetcherfpln"`
}

type MongoDBCfg struct {
	Hostip   string `mapstructure:"hostip"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type FetcherploCfg struct {
	Code string `mapstructure:"code"`
	Url  string `mapstructure:"url"`
}

type FetcherfplnCfg struct {
	Code string `mapstructure:"code"`
	Url  string `mapstructure:"url"`
}

func InitYaml() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../conf.d")
	viper.SetDefault("application.port", 8080)
	err := viper.ReadInConfig()
	if err != nil {
		panic("讀取設定檔出現錯誤，原因為：" + err.Error())
	}
	if err := viper.Unmarshal(&YamlC); err != nil {
		log.Fatal(err)
	}
	fmt.Println("application port = " + viper.GetString("application.port"))
}
