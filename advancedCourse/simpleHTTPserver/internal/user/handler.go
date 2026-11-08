package user

import (
	res "GolangCourse/advancedCourse/simpleHTTPserver/pkg"
	"GolangCourse/advancedCourse/simpleHTTPserver/pkg/middleware"
	"GolangCourse/advancedCourse/simpleHTTPserver/pkg/req"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

// UserHandlerDeps содержит зависимости HTTP-обработчика пользователей.
type UserHandlerDeps struct {
	UserRepository *UserRepository
}

// UserHandler обрабатывает HTTP-запросы к ресурсам пользователей.
type UserHandler struct {
	UserRepository *UserRepository
}

// NewUserHandler создаёт обработчик и регистрирует маршруты пользователей.
func NewUserHandler(router *http.ServeMux, deps UserHandlerDeps) {
	handler := &UserHandler{
		UserRepository: deps.UserRepository,
	}
	router.HandleFunc("POST /user", handler.Create())
	router.Handle("PATCH /user/{id}", middleware.IsAuthed(handler.Update()))
	router.HandleFunc("DELETE /user/{id}", handler.Delete())
	router.HandleFunc("GET /user/{name}", handler.GoTo())
}

// Create возвращает обработчик регистрации пользователя.
func (handler *UserHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {

		// Обработка body структуры UserCreateRequest
		body, err := req.HandleBody[UserCreateRequest](&w, request)
		if err != nil {
			return
		}

		// Создание пользователя
		user := NewUser(body.Email, body.Name, body.Password)
		// Проверка на дубль по имени
		found, err := handler.UserRepository.GetByName(user.Name)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			res.JSONRes(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if found != nil {
			res.JSONRes(w, "Пользователь с таким именем уже существует", http.StatusConflict)
			return
		}
		// Проверка на дубль по email
		found, err = handler.UserRepository.GetByEmail(user.Email)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			res.JSONRes(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if found != nil {
			res.JSONRes(w, "Пользователь с таким e-mail уже существует", http.StatusBadRequest)
			return
		}

		// Запись в базу данных новой ссылки
		createdUser, err := handler.UserRepository.Create(user)
		if err != nil {
			res.JSONRes(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Выдача ответа
		if err := res.JSONRes(w, createdUser, http.StatusCreated); err != nil {
			res.JSONRes(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}
// Update возвращает обработчик обновления пользователя.
func (handler *UserHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {}
}
// Delete возвращает обработчик удаления пользователя.
func (handler *UserHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {}
}
// GoTo возвращает обработчик получения пользователя по имени.
func (handler *UserHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {}
}
