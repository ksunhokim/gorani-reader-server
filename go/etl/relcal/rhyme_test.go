package relcal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunho/gorani-reader-server/go/etl/relcal"
)

func TestRhymeCalculatorSimple(t *testing.T) {
	a := assert.New(t)
	graph, err := relcal.Calculate("rhyme", testSet1, 0)
	a.Nil(err)

	solution := relcal.Graph{
		RelType: "rhyme",
		Vertexs: []relcal.Vertex{
			relcal.Vertex{
				WordId: 1,
				Edges: []relcal.Edge{
					relcal.Edge{
						TargetId: 2,
						Score:    3,
					},
				},
			},
			relcal.Vertex{
				WordId: 2,
				Edges: []relcal.Edge{
					relcal.Edge{
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
	graph, err := relcal.Calculate("rhyme", testSet2, 0)
	a.Nil(err)

	solution := relcal.Graph{
		RelType: "rhyme",
		Vertexs: []relcal.Vertex{
			relcal.Vertex{
				WordId: 1,
				Edges: []relcal.Edge{
					relcal.Edge{
						TargetId: 4,
						Score:    1,
					},
				},
			},
			relcal.Vertex{
				WordId: 2,
				Edges: []relcal.Edge{
					relcal.Edge{
						TargetId: 3,
						Score:    2,
					},
				},
			},
			relcal.Vertex{
				WordId: 3,
				Edges: []relcal.Edge{
					relcal.Edge{
						TargetId: 2,
						Score:    2,
					},
				},
			},
			relcal.Vertex{
				WordId: 4,
				Edges: []relcal.Edge{
					relcal.Edge{
						TargetId: 1,
						Score:    1,
					},
				},
			},
		},
	}
	a.Equal(solution, graph)

}
