package main

import (
	"context"
	"fmt"
	"github.com/peyzor/portfolio/views"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		component := views.Hello("portfolio")
		component.Render(context.Background(), w)
	})

	port := 8080
	fmt.Printf("Server is listening on port: %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf("localhost:%d", port), mux)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
