package etl

import (
	"net"

	"github.com/sunho/gorani-reader/server/pkg/gorani"
	"google.golang.org/grpc"
)

type Etl struct {
	grpc   *grpc.Server
	ln     net.Listener
	addr   string
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
