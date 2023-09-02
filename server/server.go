package server

import (
	"errors"
	"fmt"
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

func NewHttpServer(handler http.Handler) *HttpServer {
	return &HttpServer{
		server: &http.Server{
			Handler: handler,
			Addr:    ":8080",
		},
		CorsAllowedOrigins: []string{},
		prefix:             "",
		jwtSecret:          nil,
		endpoint:           "",
		host:               "localhost", // Default host
		port:               8080,        // Default port
	}
}

func (server *HttpServer) StartServer() {
	go func() {
		err := server.server.ListenAndServe()

		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server closed\n")
		} else if err != nil {
			fmt.Printf("error listening for server one: %s\n", err)
		}
	}()
}
