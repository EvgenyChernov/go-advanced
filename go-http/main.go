package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	NewHelloHandler(router)
	fmt.Println("Server is running on port 8081")
	server := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	server.ListenAndServe()
}
