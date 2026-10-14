package auth

import (
	"GolangCourse/advancedCourse/simpleHTTPserver/configs"
	res "GolangCourse/advancedCourse/simpleHTTPserver/pkg"
	req "GolangCourse/advancedCourse/simpleHTTPserver/pkg/req"
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
	return func(w http.ResponseWriter, request *http.Request) {
		//обработка body структуры LoginRequest
		_, err := req.HandleBody[LoginRequest](&w, request)
		if err != nil {
			return
		}
		//выдача ответа
		data := LoginResponse{
			Token: "123",
		}
		if err := res.JSONRes(w, data, http.StatusOK); err != nil {
			fmt.Println("Encode error: ", err)
			return
		}
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		//обработка body структуры RegisterRequest
		_, err := req.HandleBody[RegisterRequest](&w, request)
		if err != nil {
			return
		}
		//выдача ответа
		data := RegisterResponse{
			Token: "321",
			ID:    100,
		}
		if err := res.JSONRes(w, data, http.StatusCreated); err != nil {
			fmt.Println("encode error: ", err)
			return
		}
	}
}
