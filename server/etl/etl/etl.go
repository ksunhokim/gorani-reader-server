package etl

import "github.com/sunho/gorani-reader/server/pkg/gorani"

type Etl struct {
	Gorn   *gorani.Gorani
	Config Config
}

func New(gorn *gorani.Gorani, conf Config) (*Etl, error) {
	ap := &Etl{
		Gorn:   gorn,
		Config: conf,
	}
	return ap, nil
}
