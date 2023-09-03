package server

import (
	"fmt"
	"io"
	"net/http"
)

func getHello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s: got /hello request\n", ctx.Value("server"))
	io.WriteString(w, "Hello, HTTP!\n")
}

func RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/hello", getHello)
}
