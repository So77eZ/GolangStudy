package auth

import (
	"GolangCourse/advancedCourse/simpleHTTPserver/configs"
	res "GolangCourse/advancedCourse/simpleHTTPserver/pkg"
	"GolangCourse/advancedCourse/simpleHTTPserver/pkg/jwt"
	req "GolangCourse/advancedCourse/simpleHTTPserver/pkg/req"
	"errors"
	"log"
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

// NewAuthHandler подключает зависимости и регистрирует маршруты входа и регистрации.
func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}

	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

// Login возвращает обработчик входа с проверкой пароля и выдачей JWT.
func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		// Читаем и проверяем тело запроса; HandleBody сам отвечает при ошибке.
		body, err := req.HandleBody[LoginRequest](&w, request)
		if err != nil {
			return
		}

		// Сервис проверяет пароль и возвращает email найденного пользователя.
		email, err := handler.AuthService.Login(body.Email, body.Password, body.Name)
		if err != nil {
			// Неверные данные входа отделяем от внутренних ошибок сервиса.
			if errors.Is(err, ErrWrongCredetials) {
				http.Error(w, ErrWrongCredetials.Error(), http.StatusUnauthorized)
				log.Println("login error:", err)
				return
			}
			http.Error(w, "внутренняя ошибка сервера", http.StatusInternalServerError)
			log.Println("internal login error:", err)
			return
		}

		// Выпускаем токен только после успешной проверки данных входа.
		token, err := jwt.NewJWT(handler.Auth.Secret).Create(email)
		if err != nil {
			http.Error(w, "внутренняя ошибка сервера", http.StatusInternalServerError)
			log.Println("internal login error:", err)
			return
		}

		// Отправляем подписанный токен в JSON-объекте.
		data := LoginResponse{Token: token}

		res.JSONRes(w, data, http.StatusOK)
	}
}

// Register возвращает обработчик создания пользователя с последующей выдачей JWT.
func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		// Читаем и проверяем тело запроса; HandleBody сам отвечает при ошибке.
		body, err := req.HandleBody[RegisterRequest](&w, request)
		if err != nil {
			return
		}

		// Сервис сохраняет пользователя с хешем пароля и возвращает его email.
		regUserEmail, err := handler.AuthService.Register(body.Email, body.Password, body.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		// Выпускаем токен только после успешного создания пользователя.
		token, err := jwt.NewJWT(handler.Auth.Secret).Create(regUserEmail)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Формируем ответ с токеном для зарегистрированного пользователя.
		resAnswer := RegisterResponse{
			Token: token,
		}

		// При ошибке записи логируем причину: отправлять второй HTTP-ответ уже поздно.
		if err = res.JSONRes(w, resAnswer, http.StatusOK); err != nil {
			log.Println(err.Error())
			return
		}

	}
}
