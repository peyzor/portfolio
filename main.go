package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello portfolio")
	})

	port := 8080
	fmt.Printf("Server is listening on port: %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf("localhost:%d", port), mux)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
