package main

import (
	"fmt"
	"net/http"

	"app/adv-http/configs"
	"app/adv-http/internal/auth"
	"app/adv-http/internal/link"
	"app/adv-http/pkg/db"
)

func main() {

	config := configs.LoadConfig()
	_ = db.NewDB(config)
	router := http.NewServeMux()
	// hello.NewHelloHandler(router)
	auth.NewAuthHandler(router, &auth.AuthHendlerDeps{
		Config: config,
	})

	link.NewLinkHandler(router, &link.LinkHendlerDeps{})

	fmt.Println("Server is running on port 8081")
	server := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	server.ListenAndServe()
}
