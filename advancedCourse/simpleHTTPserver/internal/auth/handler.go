package auth

import (
	"GolangCourse/advancedCourse/simpleHTTPserver/configs"
	res "GolangCourse/advancedCourse/simpleHTTPserver/pkg"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
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

// mailAdress, err := mail.ParseAddress(payload.Email)//alt variant
var emailRe = regexp.MustCompile(`^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$`)

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		//Прочитать структуру входящего body
		var payload LoginRequest
		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			res.JSONRes(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if payload.Email == "" {
			fmt.Println("Введите логин")
			res.JSONRes(w, "Email is required", http.StatusBadRequest)
			return
		}

		if !emailRe.MatchString(payload.Email) {
			fmt.Println("Логин введен неправильном формате")
			res.JSONRes(w, "Wrong email format", http.StatusBadRequest)
			return
		}

		if payload.Password == "" {
			fmt.Println("Введите пароль")
			res.JSONRes(w, "Password is required", http.StatusBadRequest)
			return
		}

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
