package main

import (
	"fmt"
	"go/links-shorter/configs"
	"go/links-shorter/internal/auth"
	"go/links-shorter/internal/halthcheck"
	"go/links-shorter/internal/link"
	"go/links-shorter/pkg/db"
	"go/links-shorter/pkg/middleware"
	"net/http"
)

func main() {
	// config
	conf := configs.GetConfig()
	db := db.NewDb(conf)

	// repositories
	linkRepo := link.NewLinkRepository(db)

	// handlers
	router := http.NewServeMux()
	halthcheck.NewHalthHandler(router)
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{DbConfig: &conf.Db, AuthConfig: &conf.Auth})
	link.NewLinkHandler(router, link.LinkHandlerDeps{Repo: linkRepo})

	// middlewares
	stack := middleware.Chain(middleware.Cors, middleware.Logging)

	// server
	server := http.Server{
		Addr:    ":8080",
		Handler: stack(router), // middleware.Cors(middleware.Logging(router)),
	}

	fmt.Println("Server started on port 8080")
	server.ListenAndServe()

}
