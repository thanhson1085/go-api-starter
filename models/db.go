package models

import (
	"github.com/thanhson1085/go-api-starter/config"
	mgo "gopkg.in/mgo.v2"
)

var client *mgo.Session

func Init() {
	config := config.GetConfig()
	client, err := mgo.Dial(config.GetString("db.uri"))
	if err != nil {
		panic(err)
	}
	client.SetMode(mgo.Monotonic, true)
}

func GetDB() *mgo.Session {
	return client
}
