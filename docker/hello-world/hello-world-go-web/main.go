package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	// Create a new ServeMux
	httpMux := http.NewServeMux()

	// Register the helloHandler with the ServeMux
	httpMux.HandleFunc("/", helloHandler)

	// Start the server on port 8080
	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", httpMux); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
