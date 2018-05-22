package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sunho/gorani-reader/server/api/config"
	"github.com/sunho/gorani-reader/server/api/gorani"
	"github.com/sunho/gorani-reader/server/api/log"
)

type Server struct {
	*http.Server
	logger log.Logger
}

func (s *Server) ListenAndServe() {
	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		s.logger.Log(log.TagSystem, log.M{
			"info": "begin listening :" + s.Addr,
		})
		if err := s.Server.ListenAndServe(); err != nil {
			s.logger.Log(log.TagSystem, log.M{
				"panic": err,
			})
			panic(err)
		}
	}()

	<-stop

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	s.Shutdown(ctx)
}

func NewServer(gorn *gorani.Gorani) *Server {
	r := NewRouter(gorn)

	hs := &http.Server{
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s := &Server{
		Server: hs,
		logger: gorn.Logger,
	}

	return configureServer(s, gorn.Config)
}

func configureServer(hs *Server, conf config.Config) *Server {
	hs.Addr = conf.ApiAddress

	return hs
}
