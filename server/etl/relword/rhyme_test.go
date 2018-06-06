package relword_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunho/gorani-reader/server/etl/relword"
)

func TestRhymeCalculatorSimple(t *testing.T) {
	a := assert.New(t)
	cal := relword.RhymeCalculator{}
	graph, err := cal.Calculate(0, testSet1)
	a.Nil(err)

	solution := relword.RelGraph{
		relword.RelVertex{
			WordId: 1,
			Edges: []relword.RelEdge{
				relword.RelEdge{
					TargetId: 2,
					Score:    3,
				},
			},
		},
		relword.RelVertex{
			WordId: 2,
			Edges: []relword.RelEdge{
				relword.RelEdge{
					TargetId: 1,
					Score:    3,
				},
			},
		},
	}
	a.Equal(solution, graph)
}

func TestRhymeCalculatorComplex(t *testing.T) {
	a := assert.New(t)
	cal := relword.RhymeCalculator{}
	graph, err := cal.Calculate(0, testSet2)
	a.Nil(err)

	solution := relword.RelGraph{
		relword.RelVertex{
			WordId: 1,
			Edges: []relword.RelEdge{
				relword.RelEdge{
					TargetId: 2,
					Score:    1,
				},
				relword.RelEdge{
					TargetId: 3,
					Score:    1,
				},
			},
		},
		relword.RelVertex{
			WordId: 2,
			Edges: []relword.RelEdge{
				relword.RelEdge{
					TargetId: 1,
					Score:    1,
				},
				relword.RelEdge{
					TargetId: 3,
					Score:    2,
				},
			},
		},
		relword.RelVertex{
			WordId: 3,
			Edges: []relword.RelEdge{
				relword.RelEdge{
					TargetId: 1,
					Score:    1,
				},
				relword.RelEdge{
					TargetId: 2,
					Score:    2,
				},
			},
		},
		relword.RelVertex{
			WordId: 4,
			Edges:  []relword.RelEdge{},
		},
	}
	a.Equal(solution, graph)
}
