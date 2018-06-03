package etl

import (
	"net"

	"github.com/sunho/gorani-reader/server/pkg/gorani"
	"github.com/sunho/gorani-reader/server/pkg/log"
	"github.com/sunho/gorani-reader/server/pkg/util"
	pb "github.com/sunho/gorani-reader/server/proto/etl"
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
	e := &Etl{
		grpc:   grpc.NewServer(),
		addr:   conf.Address,
		Gorn:   gorn,
		Config: conf,
	}
	pb.RegisterETLServer(e.grpc, e)
	return e, nil
}

func (e *Etl) Addr() string {
	return e.ln.Addr().String()
}

func (e *Etl) Open() error {
	ln, err := net.Listen("tcp", e.addr)
	if err != nil {
		return err
	}

	e.ln = ln
	log.Log(log.TopicSystem, util.M{
		"msg": "begin listening " + e.addr,
	})

	go func() {
		err := e.grpc.Serve(e.ln)
		if err != nil {
			log.Log(log.TopicError, util.M{
				"msg": "grpc error ",
				"er":  err,
			})
		}
	}()

	return nil
}

func (e *Etl) Close() error {
	e.grpc.GracefulStop()
	return nil
}
