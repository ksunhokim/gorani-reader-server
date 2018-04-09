package model

import (
	"github.com/go-bongo/bongo"
)

type Unkown struct {
	bongo.DocumentBase `bson:",inline"`
	Words              []UnkownWord `json:"words"`
}

type UnkownWord struct {
	Word       string `json:"word"`
	Definition uint   `json:"def"`
	Book       string `json:"book"`
}
