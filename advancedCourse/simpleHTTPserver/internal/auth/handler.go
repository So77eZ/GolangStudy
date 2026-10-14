package auth

import (
	"GolangCourse/advancedCourse/simpleHTTPserver/configs"
	res "GolangCourse/advancedCourse/simpleHTTPserver/pkg"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
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

var validate = validator.New()

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		//Прочитать структуру входящего body
		var payload LoginRequest
		if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
			res.JSONRes(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if validate.Struct(payload) != nil {
			res.JSONRes(w, "Validation error", http.StatusBadRequest)
			return
		}

		data := LoginResponse{
			Token: "123"}
		if err := res.JSONRes(w, data, http.StatusOK); err != nil {
			fmt.Println("Encode error: ", err)
			return
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
