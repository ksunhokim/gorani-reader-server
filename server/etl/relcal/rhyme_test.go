package relcal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRhymeCalculatorSimple(t *testing.T) {
	a := assert.New(t)
	c := rhymeCalculator{}
	graph, err := c.Calculate(testSet1, 0)
	a.Nil(err)

	solution := Graph{
		Vertexs: []Vertex{
			Vertex{
				WordId: 1,
				Edges: []Edge{
					Edge{
						TargetId: 2,
						Score:    3,
					},
				},
			},
			Vertex{
				WordId: 2,
				Edges: []Edge{
					Edge{
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
	c := rhymeCalculator{}
	graph, err := c.Calculate(testSet2, 0)
	a.Nil(err)

	solution := Graph{
		Vertexs: []Vertex{
			Vertex{
				WordId: 1,
				Edges: []Edge{
					Edge{
						TargetId: 4,
						Score:    1,
					},
				},
			},
			Vertex{
				WordId: 2,
				Edges: []Edge{
					Edge{
						TargetId: 3,
						Score:    2,
					},
				},
			},
			Vertex{
				WordId: 3,
				Edges: []Edge{
					Edge{
						TargetId: 2,
						Score:    2,
					},
				},
			},
			Vertex{
				WordId: 4,
				Edges: []Edge{
					Edge{
						TargetId: 1,
						Score:    1,
					},
				},
			},
		},
	}
	a.Equal(solution, graph)

}
