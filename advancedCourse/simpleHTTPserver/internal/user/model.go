// Package user содержит модель пользователя и логику работы с ней.
package user

import (
	"math/rand/v2"

	"gorm.io/gorm"
)

// User описывает пользователя, сохраняемого в базе данных.
type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"uniqueIndex"`
	Name     string `json:"name" gorm:"uniqueIndex"`
	Password string `json:"-"`
}

// NewUser создаёт пользователя из данных регистрации.
func NewUser(email, name, password string) *User {
	user := &User{
		Email:    email,
		Name:     name,
		Password: password,
	}
	// user.GenerateHashedPassword()
	return user
}

// GenerateHashedPassword задаёт случайную строку из 10 символов; это не хеширование пароля.
func (user *User) GenerateHashedPassword() {
	user.Password = RandStringRunes(10)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()_+=-")

// RandStringRunes возвращает случайную строку длины n из letterRunes.
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	return string(b)
}
