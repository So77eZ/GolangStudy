package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()
		//Модификация обычного w для хранения полученного статус кода
		wrapper := &WrapperWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapper, r)
		log.Println("Запрос обработан за:", time.Since(start), "по пути:", r.URL.Path, "методом:", r.Method, "со статусом:", wrapper.statusCode)

	})
}
