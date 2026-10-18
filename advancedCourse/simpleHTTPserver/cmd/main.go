package main

import (
	"GolangCourse/advancedCourse/simpleHTTPserver/configs"
	"GolangCourse/advancedCourse/simpleHTTPserver/internal/auth"
	"GolangCourse/advancedCourse/simpleHTTPserver/internal/link"
	"GolangCourse/advancedCourse/simpleHTTPserver/pkg/db"
	"fmt"
	"net/http"
)

func main() {

	conf := configs.LoadConfig()
	_ = db.NewDb(conf)
	router := http.NewServeMux()

	//Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{})
	link.NewLinkHandler(router, link.LinkHandlerDeps{})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("server is listening on port: 8081")
	server.ListenAndServe()
}
