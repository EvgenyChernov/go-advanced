package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"app/adv-http/configs"
	"app/adv-http/internal/auth"
	"app/adv-http/internal/link"
	"app/adv-http/internal/user"
	"app/adv-http/pkg/db"
	"app/adv-http/pkg/middleware"
)

func main() {
	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	done := make(chan struct{})
	go func() {
		defer close(done)
		time.Sleep(3 * time.Second)
	}()

	select {
	case <-done:
		fmt.Println("done")

	case <-ctxWithTimeout.Done():
		fmt.Println("context done")
	}
}

func main2() {

	config := configs.LoadConfig()
	database := db.NewDB(config)
	router := http.NewServeMux()

	// repositories
	linkRepository := link.NewLinkRepository(database)
	userRepository := user.NewUserRepository(database)

	// services
	authService := auth.NewAuthService(userRepository)

	// handlers
	auth.NewAuthHandler(router, &auth.AuthHendlerDeps{
		Config:      config,
		AuthService: authService,
	})

	link.NewLinkHandler(router, link.LinkHendlerDeps{
		LinkRepository: linkRepository,
	})

	fmt.Println("Server is running on port 8081")
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)
	server := &http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}
	server.ListenAndServe()
}
