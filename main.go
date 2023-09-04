package main

import (
	"context"
	"fmt"

	rpc "main/rpc/server"
	"main/server"
	"net/http"
)

func httpServer() {
	//config := cmd.ParseConfigFromArgs()
	mux := http.NewServeMux()
	ctx, cancelCtx := context.WithCancel(context.Background())
	server.RegisterHandlers(mux)
	fmt.Print("Handlers Registered\n")
	httpServer := server.NewHttpServer(mux, ctx)
	httpServer.StartServer(cancelCtx)
	fmt.Print("Created HTTP Server\n")

	<-ctx.Done()
}

func main() {
	ctx, cancelCtx := context.WithCancel(context.Background())
	rpcServer := rpc.NewRPCServer()
	rpcServer.StartServer(cancelCtx)
	fmt.Print(cancelCtx, "\n")
	httpServer()

	<-ctx.Done()
}
