package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		templatePath := filepath.Join("views", "ping.html")
		tmpl, err := template.ParseFiles(templatePath)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

	port := 8080
	fmt.Printf("Server is listening on port: %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf("localhost:%d", port), mux)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
