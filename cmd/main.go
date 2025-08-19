package main

import (
	"fmt"
	"go/links-shorter/configs"
	"go/links-shorter/internal/auth"
	"go/links-shorter/internal/halthcheck"
	"go/links-shorter/pkg/db"
	"net/http"
)

func main() {

	conf := configs.GetConfig()
	// _ = conf

	_ = db.NewDb(conf)

	router := http.NewServeMux()

	halthcheck.NewHalthHandler(router)
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{DbConfig: &conf.Db, AuthConfig: &conf.Auth})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server started on port 8080")
	server.ListenAndServe()

}
