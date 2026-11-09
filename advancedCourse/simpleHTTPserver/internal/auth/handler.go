package auth

import (
	"GolangCourse/advancedCourse/simpleHTTPserver/configs"
	req "GolangCourse/advancedCourse/simpleHTTPserver/pkg/req"
	"fmt"
	"net/http"
)

// AuthHandlerDeps структура, которая содержит зависимости для AuthHandler
type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

// AuthHandler структура, которая будет обрабатывать запросы на /auth
type AuthHandler struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}

	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		//обработка body структуры LoginRequest
		body, err := req.HandleBody[LoginRequest](&w, request)
		if err != nil {
			return
		}

		email, err := handler.AuthService.Login(body.Email, body.Password, body.Name)
		if err != nil {
			fmt.Println("login error: ", err)
			return
		}
		fmt.Println("founded email:", email)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		//обработка body структуры RegisterRequest
		body, err := req.HandleBody[RegisterRequest](&w, request)
		if err != nil {
			return
		}

		if _, err = handler.AuthService.Register(body.Email, body.Password, body.Name); err != nil {
			fmt.Println("register error: ", err)
			return
		}
	}
}
