package relword

import (
	"errors"
	"sync"

	"github.com/jinzhu/gorm"
	"github.com/sunho/gorani-reader/server/pkg/dbh"
)

var (
	ErrAlreadyExist = errors.New("relword: Calculator with the same reltype already exists")
	ErrNoSuch       = errors.New("relword: No such reltype")
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

func Calculate(db *gorm.DB, reltype string, words []dbh.Word, minscore int) error {
	cal, err := calculators.get(reltype)
	if err != nil {
		return err
	}

	graph, err := cal.Calculate(words, minscore)
	if err != nil {
		return err
	}

	graph.Reltype = cal.RelType()

	err = graph.upsertToDB(db)
	if err != nil {
		return err
	}

	return nil
}
