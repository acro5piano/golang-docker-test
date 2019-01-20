package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello from Go is running at http://localhost:8000")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GET /")
		fmt.Fprintf(w, "Hello from Go")
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GET /login")
		fmt.Fprintf(w, "login from Go")
	})
	http.ListenAndServe(":8000", nil)
}
