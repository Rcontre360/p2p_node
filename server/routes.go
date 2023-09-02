package server

import (
	"fmt"
	"io"
	"net/http"
)

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/hello", getHello)
}
