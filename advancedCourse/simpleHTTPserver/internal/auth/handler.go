package auth

import (
	"GolangCourse/advancedCourse/simpleHTTPserver/configs"
	"fmt"
	"net/http"
)

// AuthHandlerDeps структура, которая содержит зависимости для AuthHandler
type AuthHandlerDeps struct {
	*configs.Config
}

// AuthHandler структура, которая будет обрабатывать запросы на /auth
type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{Config: deps.Config}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		//fmt.Println(handler.Config.Auth.Secret) // Используем конфигурацию для аутентификации
		fmt.Println("endpoint /auth/login requested")
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("endpoint /auth/register requested")
	}
}
