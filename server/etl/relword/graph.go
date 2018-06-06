package relword

type RelEdge struct {
	TargetId int32
	Score    int32
}

type RelVertex struct {
	WordId int32
	Edges  []RelEdge
}

type RelGraph []RelVertex
