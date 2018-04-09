package model

import (
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

type Unkown struct {
	bongo.DocumentBase `bson:",inline"`
	UserID             bson.ObjectId `json:"user_id"`
	Words              []UnkownWord  `json:"words"`
}

type UnkownWord struct {
	Word       string `json:"word"`
	Definition uint   `json:"def"`
	Book       string `json:"book"`
}
