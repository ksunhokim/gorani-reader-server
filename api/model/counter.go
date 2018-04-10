package model

import (
	"math/rand"

	"github.com/sirupsen/logrus"
	"github.com/sunho/engbreaker/pkg/dbs"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Counter struct {
	Name  string
	Value int
}

func NextSeq(name string) int {
	sess := dbs.MDB.Copy()
	defer sess.Close()

	change := mgo.Change{
		Upsert: true,
		Update: bson.M{
			"$inc": bson.M{
				"value": 1,
			},
		}}
	result := Counter{}
	_, err := sess.DB("").C("counters").Find(
		bson.M{
			"name": name,
		}).Apply(change, &result)

	if err != nil {
		logrus.Error("Model counter error:", name, "   ", err)
		return rand.Intn(5000000)
	}
	return result.Value
}
