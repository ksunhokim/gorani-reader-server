package relword

import (
	"errors"
	"sync"

	"github.com/sunho/gorani-reader/server/pkg/dbh"
)

var (
	ErrAlreadyExist = errors.New("relword: Calculator with the same reltype already exists")
	ErrNoSuch       = errors.New("relword: No such reltype")
)

type Calculator interface {
	Calculate(minscore int, words []dbh.Word) (Graph, error)
	RelType() string
}

type calculatorSlice struct {
	mu   sync.RWMutex
	cals map[string]Calculator
}

func (cs *calculatorSlice) get(typ string) (Calculator, error) {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	if cal, ok := cs.cals[typ]; ok {
		return cal, nil
	}
	return nil, ErrNoSuch
}

func (cs *calculatorSlice) add(cal Calculator) error {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	if _, ok := cs.cals[cal.RelType()]; ok {
		return ErrAlreadyExist
	}
	cs.cals[cal.RelType()] = cal
	return nil
}

var calculators calculatorSlice

func init() {
	calculators = calculatorSlice{
		cals: make(map[string]Calculator),
	}
}

func RegisterCalculator(cal Calculator) error {
	return calculators.add(cal)
}

func Calculate(reltype string, words []dbh.Word, minscore int) (graph Graph, err error) {
	cal, err := calculators.get(reltype)
	if err != nil {
		return
	}

	graph, err = cal.Calculate(minscore, words)
	if err != nil {
		return
	}

	graph.Reltype = cal.RelType()

	return
}
