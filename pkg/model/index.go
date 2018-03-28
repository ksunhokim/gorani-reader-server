package model

import (
	"github.com/sunho/engbreaker/pkg/dbs"
	mgo "gopkg.in/mgo.v2"
)

func MigrateIndex() {
	dbs.MDB.Collection("wordbooks").Collection().EnsureIndex(
		mgo.Index{
			Key:    []string{"name", "userid"},
			Unique: true,
		},
	)
	dbs.MDB.Collection("users").Collection().EnsureIndex(
		mgo.Index{
			Key:    []string{"authprovider", "authid"},
			Unique: true,
		},
	)
}
