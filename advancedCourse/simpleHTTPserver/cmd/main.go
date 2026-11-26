package main

import (
	"GolangCourse/advancedCourse/simpleHTTPserver/configs"
	"GolangCourse/advancedCourse/simpleHTTPserver/internal/auth"
	"GolangCourse/advancedCourse/simpleHTTPserver/internal/link"
	"GolangCourse/advancedCourse/simpleHTTPserver/internal/stat"
	"GolangCourse/advancedCourse/simpleHTTPserver/internal/user"
	"GolangCourse/advancedCourse/simpleHTTPserver/pkg/db"
	"GolangCourse/advancedCourse/simpleHTTPserver/pkg/event"
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
	// Обработчик редиректов и сервис статистики используют одну шину событий.
	eventBus := event.NewEventBus()

	//Repositories
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)
	statRepository := stat.NewStatRepository(db)

	//Services
	authService := auth.NewAuthService(userRepository)
	statService := stat.NewStatService(&stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepository,
	})

	//Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService})

	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		EventBus:       eventBus,
		Config:         conf})

	user.NewUserHandler(router, user.UserHandlerDeps{
		UserRepository: userRepository,
		Config:         conf})

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

	// Запускаем один цикл чтения событий на весь сервер.
	go statService.AddClick()

	error := server.ListenAndServe()
	if error != nil {
		panic(error)
	}
}
