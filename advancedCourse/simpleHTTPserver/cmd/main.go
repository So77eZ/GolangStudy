package main

import (
	"GolangCourse/advancedCourse/simpleHTTPserver/configs"
	"GolangCourse/advancedCourse/simpleHTTPserver/internal/auth"
	"GolangCourse/advancedCourse/simpleHTTPserver/pkg/db"
	"fmt"
	"net/http"
)

func main() {

	conf := configs.LoadConfig()
	_ = db.NewDb(conf)
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{Config: conf})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("server is listening on port: 8081")
	server.ListenAndServe()
}
