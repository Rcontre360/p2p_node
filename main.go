package main

import (
	"context"
	"main/rpc"
	"main/server"
	"net/http"
)

func httpServer() {
	//config := cmd.ParseConfigFromArgs()
	mux := http.NewServeMux()
	ctx, cancelCtx := context.WithCancel(context.Background())

	server.RegisterHandlers(mux)

	httpServer := server.NewHttpServer(mux, ctx)
	httpServer.StartServer(cancelCtx)

	<-ctx.Done()
}

func main() {
	ctx, cancelCtx := context.WithCancel(context.Background())
	rpcServer := rpc.NewRPCServer()

	rpcServer.StartServer(cancelCtx)

	<-ctx.Done()
}
