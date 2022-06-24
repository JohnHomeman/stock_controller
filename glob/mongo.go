package glob

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Client

type MongoCfg struct {
	Host string
}

func NewMongoCfg() *MongoCfg {
	return &MongoCfg{
		Host: YamlC.MongoDBCfg.Hostip,
	}
}

func NewMongoServer(config *MongoCfg) *mongo.Client {
	urlMongo := "mongodb://" + config.Host
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(urlMongo))
	if err != nil {
		panic(err)
	}

	return client
}
