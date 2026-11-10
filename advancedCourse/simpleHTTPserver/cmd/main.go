package main

import (
	"GolangCourse/advancedCourse/simpleHTTPserver/configs"
	"GolangCourse/advancedCourse/simpleHTTPserver/internal/auth"
	"GolangCourse/advancedCourse/simpleHTTPserver/internal/link"
	"GolangCourse/advancedCourse/simpleHTTPserver/internal/user"
	"GolangCourse/advancedCourse/simpleHTTPserver/pkg/db"
	"GolangCourse/advancedCourse/simpleHTTPserver/pkg/middleware"
	"fmt"
	"net/http"
)

func main() {

	conf := configs.LoadConfig()
	if conf.Auth.Secret == "" || conf.Auth.Secret == " " {
		panic("Не задан токен для конфига")
	}

	db := db.NewDb(conf)
	router := http.NewServeMux()

	//Repositories
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)

	//Services
	authService := auth.NewAuthService(userRepository)

	//Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService})
	link.NewLinkHandler(router, link.LinkHandlerDeps{LinkRepository: linkRepository})
	user.NewUserHandler(router, user.UserHandlerDeps{UserRepository: userRepository})

	//Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}

	fmt.Println("server is listening on port: 8081")
	error := server.ListenAndServe()
	if error != nil {
		panic(error)
	}
}
