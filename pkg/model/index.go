package model

import (
	"github.com/sunho/engbreaker/pkg/dbs"
	mgo "gopkg.in/mgo.v2"
)

func MigrateIndex() {
	index := mgo.Index{
		Key:    []string{"name"},
		Unique: true,
	}
	dbs.MDB.Collection("wordbooks").Collection().EnsureIndex(index)
}
