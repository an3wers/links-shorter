package main

import (
	"fmt"
	"go/links-shorter/configs"
	"go/links-shorter/internal/auth"
	"go/links-shorter/internal/halthcheck"
	"go/links-shorter/internal/link"
	"go/links-shorter/internal/stat"
	"go/links-shorter/internal/user"
	"go/links-shorter/pkg/db"
	"go/links-shorter/pkg/event"
	"go/links-shorter/pkg/middleware"
	"net/http"
)

func main() {
	// config
	conf := configs.GetConfig()

	// database
	db := db.NewDb(conf)

	// EventBus
	eventBus := event.NewEventBus()

	// repositories
	linkRepo := link.NewLinkRepository(db)
	userRepo := user.NewUserRepository(db)
	statRepo := stat.NewStatRepository(db)

	// services
	authService := auth.NewAuthService((userRepo))
	statService := stat.NewStatService(stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepo,
	})

	// listen events
	go statService.ListenEvents()

	// create router (Mux)
	router := http.NewServeMux()

	// handlers
	halthcheck.NewHalthHandler(router)
	auth.NewAuthHandler(
		router, auth.AuthHandlerDeps{DbConfig: &conf.Db,
			AuthConfig:  &conf.Auth,
			AuthService: authService,
		})

	link.NewLinkHandler(router, link.LinkHandlerDeps{
		Repo:     linkRepo,
		EventBus: eventBus,
		Config:   conf,
	})

	// middlewares
	stack := middleware.Chain(middleware.Cors, middleware.Logging)

	// server
	server := http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}

	// start
	fmt.Println("Server started on port 8080")
	server.ListenAndServe()

}
