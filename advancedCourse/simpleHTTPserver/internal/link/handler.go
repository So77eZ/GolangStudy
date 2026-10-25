package link

import (
	res "GolangCourse/advancedCourse/simpleHTTPserver/pkg"
	"GolangCourse/advancedCourse/simpleHTTPserver/pkg/req"
	"fmt"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

// LinkHandlerDeps структура, которая содержит зависимости для LinkHandler
type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
}

// LinkHandler структура, которая будет обрабатывать запросы на /link CRUD
type LinkHandler struct {
	LinkRepository *LinkRepository
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
	}
	router.HandleFunc("POST /link", handler.Create())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
	router.HandleFunc("GET /{hash}", handler.GoTo())
}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {

		// Обработка body структуры LinkCreateRequest
		body, err := req.HandleBody[LinkCreateRequest](&w, request)
		if err != nil {
			return
		}

		// Создание хэша и проверка его на уникальность
		link := NewLink(body.URL)
		var triesCounter = 0
		for {

			if foundLink, _ := handler.LinkRepository.GetByHash(link.Hash); foundLink == nil {
				break
			}
			link.GenerateHash()
			triesCounter++
			if triesCounter >= 10 {
				res.JSONRes(w, "Could not generate unique hash", http.StatusInternalServerError)
				return
			}
		}
		// Запись в базу данных новой ссылки
		createdLink, err := handler.LinkRepository.Create(link)
		if err != nil {
			res.JSONRes(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Выдача ответа
		if err := res.JSONRes(w, createdLink, http.StatusCreated); err != nil {
			res.JSONRes(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
func (handler *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hashURl := r.PathValue("hash")
		link, err := handler.LinkRepository.GetByHash(hashURl)
		if err != nil {
			res.JSONRes(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Redirect(w, r, link.URL, http.StatusTemporaryRedirect)

	}
}
func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Обработка body структуры LinkUpdateRequest
		body, err := req.HandleBody[LinkUpdateRequest](&w, r)
		if err != nil {
			return
		}
		// Собираем пришедший id записи и конвертируем его в формат gorm модели
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 64)
		if err != nil {
			res.JSONRes(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Проверка существования записи
		_, err = handler.LinkRepository.GetByID(uint(id))
		if err != nil {
			res.JSONRes(w, err.Error(), http.StatusNotFound)
			return
		}
		// Само обновление записи в БД полученными параметрами
		link, err := handler.LinkRepository.Update(&Link{
			Model: gorm.Model{ID: uint(id)},
			URL:   body.URL,
			Hash:  body.Hash,
		})
		if err != nil {
			res.JSONRes(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Выдача ответа
		if err := res.JSONRes(w, link, http.StatusOK); err != nil {
			fmt.Println("Encode error: ", err)
			return
		}

	}
}
func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Собираем пришедший id записи и конвертируем его в формат gorm модели
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 64)
		if err != nil {
			res.JSONRes(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Проверка существования записи
		_, err = handler.LinkRepository.GetByID(uint(id))
		if err != nil {
			res.JSONRes(w, err.Error(), http.StatusNotFound)
			return
		}
		// Само удаление записи в БД по переданной ссылке
		if err = handler.LinkRepository.Delete(uint(id)); err != nil {
			res.JSONRes(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Выдача ответа
		if err := res.JSONRes(w, nil, http.StatusNoContent); err != nil {
			return
		}
	}
}
