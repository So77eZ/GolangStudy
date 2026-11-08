package auth

import (
	"GolangCourse/advancedCourse/simpleHTTPserver/internal/user"
	"errors"

	"gorm.io/gorm"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (service *AuthService) Register(email, password, name string) (string, error) {

	// Проверка на дубль по имени
	existedUser, err := service.UserRepository.GetByName(name)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", errors.New("запись не найдена")
	}
	if existedUser != nil {
		return "", errors.New(ErrUserNameExists)
	}

	// Проверка на дубль по email
	existedUser, err = service.UserRepository.GetByEmail(email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", errors.New("запись не найдена")
	}
	if existedUser != nil {
		return "", errors.New(ErrUserEmailExists)
	}

	// Создание нового пользователя
	user := &user.User{
		Email:    email,
		Password: password,
		Name:     name,
	}

	if _, err = service.UserRepository.Create(user); err != nil {
		return "", errors.New(ErrRegistrationErr)
	}
	return user.Email, nil
}
