package model

type Unkown struct {
	Words []UnkownWord `json:"words"`
}

type UnkownWord struct {
	Word       string `json:"word"`
	Definition uint   `json:"def"`
	Book       string `json:"book"`
}
