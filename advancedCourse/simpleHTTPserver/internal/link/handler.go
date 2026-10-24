package link

import (
	res "GolangCourse/advancedCourse/simpleHTTPserver/pkg"
	"GolangCourse/advancedCourse/simpleHTTPserver/pkg/req"
	"fmt"
	"net/http"
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
		for {

			if foundLink, _ := handler.LinkRepository.GetByHash(link.Hash); foundLink == nil {
				break
			}
			link.GenerateHash()
		}
		// Запись в базу данных новой ссылки
		createdLink, err := handler.LinkRepository.Create(link)
		if err != nil {
			res.JSONRes(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Выдача ответа
		if err := res.JSONRes(w, createdLink, http.StatusCreated); err != nil {
			fmt.Println("Encode error: ", err)
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
		fmt.Println("u'r in UPDATE link page")
	}
}
func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Println(id)
	}
}
