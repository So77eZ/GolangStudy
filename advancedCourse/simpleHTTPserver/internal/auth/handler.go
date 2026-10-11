package auth

import (
	"GolangCourse/advancedCourse/simpleHTTPserver/configs"
	res "GolangCourse/advancedCourse/simpleHTTPserver/pkg"
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
		fmt.Println("endpoint /auth/login requested")
		data := LoginResponse{
			Token: "123"}
		if err := res.JSONRes(w, data, http.StatusOK); err != nil {
			fmt.Println("encode error: ", err)
		}
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("endpoint /auth/register requested")
		data := RegisterResponse{
			Token: "333",
			Email: "ja227@mail.com",
			ID:    12399191,
			Name:  "Slava",
		}
		if err := res.JSONRes(w, data, http.StatusCreated); err != nil {
			fmt.Println("encode error: ", err)
		}
	}
}
