package main

import (
	"GolangCourse/advancedCourse/simpleHTTPserver/configs"
	"GolangCourse/advancedCourse/simpleHTTPserver/internal/auth"
	"fmt"
	"net/http"
)

func main() {

	conf := configs.LoadConfig()
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{Config: conf})

	// /auth/login
	// /auth/register

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("server is listening on port: 8081")
	server.ListenAndServe()
}
