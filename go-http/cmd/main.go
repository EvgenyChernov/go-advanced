package main

import (
	"fmt"
	"net/http"

	"app/adv-http/configs"
	"app/adv-http/internal/auth"
	"app/adv-http/internal/link"
	"app/adv-http/internal/stat"
	"app/adv-http/internal/user"
	"app/adv-http/pkg/db"
	"app/adv-http/pkg/event"
	"app/adv-http/pkg/middleware"
)

func App() http.Handler {
	config := configs.LoadConfig()
	database := db.NewDB(config)
	statRepository := stat.NewStatRepository(database)
	router := http.NewServeMux()
	eventBus := event.NewEventBus()
	// repositories
	linkRepository := link.NewLinkRepository(database)
	userRepository := user.NewUserRepository(database)

	// services
	authService := auth.NewAuthService(userRepository)
	statService := stat.NewStatService(stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepository,
	})
	// handlers
	auth.NewAuthHandler(router, &auth.AuthHendlerDeps{
		Config:      config,
		AuthService: authService,
	})

	link.NewLinkHandler(router, link.LinkHendlerDeps{
		LinkRepository: linkRepository,
		Config:         config,
		EventBus:       eventBus,
	})

	stat.NewStatHandler(router, stat.StatHandlerDeps{
		StatRepository: statRepository,
		Config:         config,
	})

	fmt.Println("Server is running on port 8081")
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)
	go statService.AddClick()
	return stack(router)
}

func main() {
	app := App()
	server := &http.Server{
		Addr:    ":8081",
		Handler: app,
	}

	server.ListenAndServe()
}
