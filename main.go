package main

import (
	"context"
	"main/server"
	"net/http"
)

func main() {
	//config := cmd.ParseConfigFromArgs()
	mux := http.NewServeMux()
	ctx, cancelCtx := context.WithCancel(context.Background())

	server.RegisterHandlers(mux)

	httpServer := server.NewHttpServer(mux, ctx)
	httpServer.StartServer(cancelCtx)

	<-ctx.Done()
}
