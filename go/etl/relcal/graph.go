package relcal

type Edge struct {
	TargetId int
	Score    int
}

type Vertex struct {
	WordId int
	Edges  []Edge
}

type Graph struct {
	Vertexs []Vertex
	RelType string
}
