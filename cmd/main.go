package main

import (
	"fmt"
	"go/links-shorter/configs"
	"go/links-shorter/internal/auth"
	"go/links-shorter/internal/halthcheck"
	"go/links-shorter/internal/link"
	"go/links-shorter/pkg/db"
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

	// server
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server started on port 8080")
	server.ListenAndServe()

}
