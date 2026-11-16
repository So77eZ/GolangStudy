package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

// JWT хранит секрет для подписи токенов.
type JWT struct {
	Secret string
}

type JWTData struct {
	Email string
}

// NewJWT создаёт объект для выпуска токенов с указанным секретом.
func NewJWT(secret string) *JWT {
	return &JWT{
		Secret: secret,
	}
}

// Create выпускает подписанный алгоритмом HS256 токен с email пользователя.
func (j *JWT) Create(data JWTData) (string, error) {
	// Claims доступны для чтения: подпись защищает их от изменения, а не скрывает.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": data.Email,
	})
	// Подписываем токен секретом и получаем строку для передачи клиенту.
	s, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return s, nil
}

func (j *JWT) Parse(token string) (bool, *JWTData) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return false, nil
	}
	email := t.Claims.(jwt.MapClaims)["email"]
	return t.Valid, &JWTData{
		Email: email.(string),
	}
}
