package main

import (
	"fmt"
	"main/cmd"
	"net/http"
)

func main() {
	config := cmd.ParseConfigFromArgs()

	fmt.Println("config: ", config)
	// Define a handler function for handling incoming HTTP requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	// Start the server on port 8080
	fmt.Println("Server listening on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
