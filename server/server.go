package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
)

type HttpServer struct {
	server *http.Server

	CorsAllowedOrigins []string
	jwtSecret          []byte // optional JWT secret
	prefix             string // path prefix on which to mount http handler

	endpoint string
	host     string
	port     int
}

func NewHttpServer(handler http.Handler, ctx context.Context) *HttpServer {

	return &HttpServer{
		server: &http.Server{
			Handler: handler,
			Addr:    ":8080",
			BaseContext: func(l net.Listener) context.Context {
				ctx = context.WithValue(ctx, "rafael", l.Addr().String())
				return ctx
			},
		},
		CorsAllowedOrigins: []string{},
		prefix:             "",
		jwtSecret:          nil,
		endpoint:           "",
		host:               "localhost", // Default host
		port:               8080,        // Default port
	}
}

func (server *HttpServer) StartServer(cancelCtx context.CancelFunc) {
	go func() {
		err := server.server.ListenAndServe()

		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server closed\n")
		} else if err != nil {
			fmt.Printf("error listening for server one: %s\n", err)
		}

		cancelCtx()
	}()
}
