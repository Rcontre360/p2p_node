package main

import (
	"context"
	"fmt"

	"main/cmd"
	rpc "main/rpc/server"
	"main/server"
	"net/http"
)

func httpServer() {
	config := cmd.ParseConfigFromArgs()
	mux := http.NewServeMux()
	ctx, cancelCtx := context.WithCancel(context.Background())
	server.RegisterHandlers(mux)
	fmt.Print("Handlers Registered\n")
	httpServer := server.NewHttpServer(mux, ctx, config)
	httpServer.StartServer(cancelCtx)
	fmt.Print("Created HTTP Server\n")

	<-ctx.Done()
}

func main() {
	config := cmd.ParseRPCConfigFromArgs()

	ctx, cancelCtx := context.WithCancel(context.Background())
	rpcServer := rpc.NewRPCServer(config)
	rpcServer.StartServer(cancelCtx)

	<-ctx.Done()
}
