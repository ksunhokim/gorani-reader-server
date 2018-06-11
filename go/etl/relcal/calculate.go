package relcal

import (
	"errors"
	"sync"

	"github.com/sunho/gorani-reader-server/go/pkg/dbh"
)

var (
	ErrAlreadyExist = errors.New("relcal: Calculator with the same reltype already exists")
	ErrNoSuch       = errors.New("relcal: No such reltype")
)

type calculator interface {
	Calculate(words []dbh.Word, minscore int) (Graph, error)
	RelType() string
}

type calculatorSlice struct {
	mu   sync.RWMutex
	cals map[string]calculator
}

func (cs *calculatorSlice) get(typ string) (calculator, error) {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	if cal, ok := cs.cals[typ]; ok {
		return cal, nil
	}
	return nil, ErrNoSuch
}

func (cs *calculatorSlice) add(cal calculator) error {
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
		cals: make(map[string]calculator),
	}
}

func Calculate(reltype string, words []dbh.Word, minscore int) (graph Graph, err error) {
	cal, err := calculators.get(reltype)
	if err != nil {
		return
	}

	graph, err = cal.Calculate(words, minscore)
	if err != nil {
		return
	}

	graph.RelType = cal.RelType()

	return
}
