package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World 2!")
	})

	fmt.Println("Server running on port 3000")
	http.ListenAndServe(":3000", nil)
}