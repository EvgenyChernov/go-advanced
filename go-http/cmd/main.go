package main

import (
	"fmt"
	"net/http"

	"app/adv-http/configs"
	"app/adv-http/internal/auth"
	"app/adv-http/internal/link"
	"app/adv-http/pkg/db"
	"app/adv-http/pkg/middleware"
)

func main() {

	config := configs.LoadConfig()
	database := db.NewDB(config)
	router := http.NewServeMux()

	// repositories
	linkRepository := link.NewLinkRepository(database)

	// hello.NewHelloHandler(router)
	auth.NewAuthHandler(router, &auth.AuthHendlerDeps{
		Config: config,
	})

	link.NewLinkHandler(router, link.LinkHendlerDeps{
		LinkRepository: linkRepository,
	})

	fmt.Println("Server is running on port 8081")
	server := &http.Server{
		Addr:    ":8081",
		Handler: middleware.Logging(router),
	}
	server.ListenAndServe()
}
