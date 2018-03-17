package api

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/sunho/engbreaker/pkg/config"
	"github.com/sunho/engbreaker/pkg/middlewares"
	"github.com/sunho/engbreaker/pkg/util"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	route   *mux.Router
	httpSrv *http.Server
}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{}
}

func (h *HTTPServer) Start() error {
	h.registerRoutes()

	addr := config.GetString(config.LISTENADDR)
	handler := middlewares.Log(h.route)
	h.httpSrv = &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	var err error
	protocol := config.GetString(config.PROTOCOL)
	switch protocol {
	case "HTTPS":
		logrus.Info("api server started as https server")
		err = h.listenAndServeTLS(config.GetString(config.CERTFILE),
			config.GetString(config.KEYFILE))
		if err == http.ErrServerClosed {
			logrus.Info("api server was shutdown gracefully")
			return nil
		}
	case "HTTP":
		fallthrough
	default:
		logrus.Info("api server started as http server")
		err = h.httpSrv.ListenAndServe()
		if err == http.ErrServerClosed {
			logrus.Info("api server was shutdown gracefully")
			return nil
		}
	}
	return err
}

func (h *HTTPServer) listenAndServeTLS(cert, key string) error {
	if cert == "" {
		return fmt.Errorf(`%s is empty`, config.CERTFILE)
	}
	if key == "" {
		return fmt.Errorf(`%s is empty`, config.KEYFILE)
	}

	if !util.FileExist(cert) {
		return fmt.Errorf(`cannot find SSL cert at %s from %s`, cert, config.CERTFILE)
	}
	if !util.FileExist(key) {
		return fmt.Errorf(`cannot find SSL key at %s from %s`, key, config.KEYFILE)
	}

	//https://github.com/denji/golang-tls
	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}
	h.httpSrv.TLSConfig = cfg
	h.httpSrv.TLSNextProto = make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0)
	h.httpSrv.Handler = middlewares.SSL(h.httpSrv.Handler)
	return h.httpSrv.ListenAndServeTLS(cert, key)
}
