package middleware

import (
	"GolangCourse/advancedCourse/simpleHTTPserver/configs"
	"GolangCourse/advancedCourse/simpleHTTPserver/pkg/jwt"
	"fmt"
	"net/http"
	"strings"
)

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
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
		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)
		fmt.Println(data.Email)
		fmt.Println(isValid)
		next.ServeHTTP(w, r)
	})
}
