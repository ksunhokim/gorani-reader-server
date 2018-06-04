package service

import (
	"net"

	"github.com/sunho/gorani-reader/server/etl/etl"
	"github.com/sunho/gorani-reader/server/pkg/log"
	"github.com/sunho/gorani-reader/server/pkg/util"
	pb "github.com/sunho/gorani-reader/server/proto/etl"
	"google.golang.org/grpc"
)

type Service struct {
	pb.ETLServiceServer
	e    *etl.Etl
	grpc *grpc.Server
	ln   net.Listener
}

func New(e *etl.Etl) *Service {
	serv := &Service{
		e:    e,
		grpc: grpc.NewServer(),
	}
	pb.RegisterETLServiceServer(serv.grpc, serv)
	return serv
}

func (s *Service) Addr() string {
	return s.ln.Addr().String()
}

func (s *Service) Open() error {
	ln, err := net.Listen("tcp", s.e.Config.Address)
	if err != nil {
		return err
	}

	s.ln = ln
	log.Log(log.TopicSystem, util.M{
		"msg": "begin listening " + s.e.Config.Address,
	})

	go func() {
		err := s.grpc.Serve(s.ln)
		if err != nil {
			log.Log(log.TopicError, util.M{
				"msg": "grpc error ",
				"er":  err,
			})
		}
	}()

	return nil
}

func (s *Service) Close() error {
	s.grpc.GracefulStop()
	return nil
}
