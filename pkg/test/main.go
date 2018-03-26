package main

import (
	"fmt"

	"github.com/sunho/engbreaker/pkg/dbs"
	"github.com/sunho/engbreaker/pkg/model"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	dbs.Init()
	arr := []*model.User{}
	model.Get(&arr, bson.M{})
	fmt.Println(arr)
}
