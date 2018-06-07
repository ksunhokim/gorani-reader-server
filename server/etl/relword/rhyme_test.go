package relword_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunho/gorani-reader/server/etl/relword"
)

func TestRhymeCalculatorSimple(t *testing.T) {
	a := assert.New(t)
	graph, err := relword.Calculate("rhyme", testSet1, 0)
	a.Nil(err)

	solution := relword.Graph{
		Reltype: "rhyme",
		Vertexs: []relword.Vertex{
			relword.Vertex{
				WordId: 1,
				Edges: []relword.Edge{
					relword.Edge{
						TargetId: 2,
						Score:    3,
					},
				},
			},
			relword.Vertex{
				WordId: 2,
				Edges: []relword.Edge{
					relword.Edge{
						TargetId: 1,
						Score:    3,
					},
				},
			},
		},
	}
	a.Equal(solution, graph)
}

func TestRhymeCalculatorComplex(t *testing.T) {
	a := assert.New(t)
	graph, err := relword.Calculate("rhyme", testSet2, 0)
	a.Nil(err)

	solution := relword.Graph{
		Reltype: "rhyme",
		Vertexs: []relword.Vertex{
			relword.Vertex{
				WordId: 1,
				Edges: []relword.Edge{
					relword.Edge{
						TargetId: 4,
						Score:    1,
					},
				},
			},
			relword.Vertex{
				WordId: 2,
				Edges: []relword.Edge{
					relword.Edge{
						TargetId: 3,
						Score:    2,
					},
				},
			},
			relword.Vertex{
				WordId: 3,
				Edges: []relword.Edge{
					relword.Edge{
						TargetId: 2,
						Score:    2,
					},
				},
			},
			relword.Vertex{
				WordId: 4,
				Edges: []relword.Edge{
					relword.Edge{
						TargetId: 1,
						Score:    1,
					},
				},
			},
		},
	}
	a.Equal(solution, graph)

}
