package main

import (
	"fmt"
	"net/http"

	"app/adv-http/configs"
	"app/adv-http/internal/hello"
)

func main() {

	config := configs.LoadConfig()
	router := http.NewServeMux()
	hello.NewHelloHandler(router)
	fmt.Println("Server is running on port 8081")
	server := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	server.ListenAndServe()
}
