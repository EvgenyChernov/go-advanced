package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("Server is running on port 8081")
	http.ListenAndServe(":8081", nil)
}
