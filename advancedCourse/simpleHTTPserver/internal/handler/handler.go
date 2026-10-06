package handler

import (
	"fmt"
	"net/http"
)

// func hello(w http.ResponseWriter, req *http.Request) {
// 	fmt.Println("Server says hello")
// }

// HelloHandler структура, которая будет обрабатывать запросы на /hello
type HelloHandler struct{}

// NewHelloHandler создает новый обработчик для маршрута /hello и регистрирует его в переданном роутере
func NewHelloHandler(router *http.ServeMux) {
	handler := &HelloHandler{}
	router.HandleFunc("/hello", handler.Hello())
}

// Hello метод, который возвращает саму
func (handler *HelloHandler) Hello() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Servers says hello")
	}

}
