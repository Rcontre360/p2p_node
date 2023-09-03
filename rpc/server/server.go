// Package main implements a server for Greeter service.
package rpc

import (
	"context"
	"fmt"
	"log"
	"main/rpc/definitions"
	"net"

	"google.golang.org/grpc"
)

// server is used to implement helloworld.GreeterServer.
type RPCServer struct {
	definitions.UnimplementedP2PServer
	server *grpc.Server

	CorsAllowedOrigins []string
	jwtSecret          []byte // optional JWT secret
	prefix             string // path prefix on which to mount http handler

	endpoint string
	host     string
	port     int
}

func NewRPCServer() *RPCServer {
	return &RPCServer{
		server:             grpc.NewServer(),
		CorsAllowedOrigins: []string{},
		prefix:             "",
		jwtSecret:          nil,
		endpoint:           "",
		host:               "localhost", // Default host
		port:               8080,        // Default port
	}
}

func (server *RPCServer) StartServer(cancelCtx context.CancelFunc) {
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", server.port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()
		definitions.RegisterP2PServer(grpcServer, server)

		if err := server.server.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}

		cancelCtx()
	}()
}
