package model

import (
	"math/rand"

	"github.com/sirupsen/logrus"
	"github.com/sunho/gorani-reader/pkg/dbs"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Counter struct {
	Name  string `bson:"name"`
	Value int    `bson:"value"`
}

func NextSeq(name string) int {
	sess := dbs.MDB.Copy()
	defer sess.Close()

	// increase or insert
	change := mgo.Change{
		ReturnNew: true,
		Upsert:    true,
		Update: bson.M{
			"$inc": bson.M{
				"value": 1,
			},
		}}

	counter := Counter{}
	_, err := sess.DB("").C("counters").Find(
		bson.M{
			"name": name,
		}).Apply(change, &counter)

	if err != nil {
		// should not occur
		logrus.Error("Model counter error:", name, "   ", err)
		return rand.Intn(5000000)
	}

	return counter.Value
}
