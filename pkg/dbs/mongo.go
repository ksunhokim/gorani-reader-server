package dbs

import (
	"github.com/sirupsen/logrus"
	"github.com/sunho/gorani-reader/pkg/config"
	mgo "gopkg.in/mgo.v2"
)

var MDB *mgo.Session

func initMongo() {
	addr := config.GetString("MONGO_ADDR", "localhost")
	db := config.GetString("MONGO_DB", "bongotest")
	info := &mgo.DialInfo{
		Addrs:    []string{addr},
		Database: db,
	}
	sess, err := mgo.DialWithInfo(info)
	if err != nil {
		logrus.Panic(err)
	}
	MDB = sess
}
