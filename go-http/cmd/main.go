package main

import (
	"fmt"
	"net/http"

	"app/adv-http/configs"
	"app/adv-http/internal/auth"
)

func main() {

	config := configs.LoadConfig()
	router := http.NewServeMux()
	// hello.NewHelloHandler(router)
	auth.NewAuthHandler(router, &auth.AuthHendlerDeps{
		Config: config,
	})
	fmt.Println("Server is running on port 8081")
	server := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	server.ListenAndServe()
}
