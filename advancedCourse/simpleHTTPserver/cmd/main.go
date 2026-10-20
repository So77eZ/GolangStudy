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
	db := db.NewDb(conf)
	router := http.NewServeMux()

	//Repositories
	linkRepository := link.NewLinkRepository(db)

	//Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{Config: conf})
	link.NewLinkHandler(router, link.LinkHandlerDeps{LinkRepository: linkRepository})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("server is listening on port: 8081")
	error := server.ListenAndServe()
	if error != nil {
		panic(error)
	}
}
