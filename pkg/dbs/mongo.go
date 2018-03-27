package dbs

import (
	"github.com/sirupsen/logrus"
	"github.com/sunho/bongo"
	"github.com/sunho/engbreaker/pkg/config"
)

var MDB *bongo.Connection

func initMongo() {
	addr := config.GetString("MONGO_ADDR", "localhost")
	db := config.GetString("MONGO_DB", "bongotest")
	tdb, err := bongo.Connect(&bongo.Config{
		ConnectionString: addr,
		Database:         db,
	})
	if err != nil {
		logrus.Panic(err)
	}
	MDB = tdb
}
