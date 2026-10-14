package req

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// Валидация структуры
func IsValid[T any](payload T) error {
	err := validate.Struct(payload)
	return err
}
