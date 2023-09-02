package main

import (
	"fmt"
	"main/cmd"
	"main/server"
	"net/http"
)

func main() {
	config := cmd.ParseConfigFromArgs()
	mux := http.NewServeMux()

	fmt.Println("config: ", config)
	server.RegisterHandlers(mux)

	httpServer := server.NewHttpServer(mux)
	httpServer.StartServer()

}
