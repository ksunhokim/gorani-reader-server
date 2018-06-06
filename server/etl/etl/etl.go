package etl

import (
	"github.com/sunho/gorani-reader/server/pkg/gorani"
)

type Etl struct {
	*gorani.Gorani
	Config Config
}

func New(gorn *gorani.Gorani, conf Config) (*Etl, error) {
	e := &Etl{
		Gorani: gorn,
		Config: conf,
	}
	e.Config.Config = gorn.Config

	return e, nil
}
