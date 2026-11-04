package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func IsAuthed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("Authorization")
		if bearer == "" {
			return //_, fmt.Errorf("Нет bearer токена")
		}
		token, ok := strings.CutPrefix(bearer, "Bearer ")
		if !ok {
			return //fmt.Errorf("неверный формат токена: ожидается 'Bearer <token>'")
		}
		if token == "" {
			return //fmt.Errorf("токен отсутствует после префикса")
		}
		fmt.Println(token)
		next.ServeHTTP(w, r)
	})
}
