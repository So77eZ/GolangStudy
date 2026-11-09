package auth

import (
	"GolangCourse/advancedCourse/simpleHTTPserver/internal/user"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// AuthService выполняет регистрацию и проверку данных для входа через репозиторий пользователей.
type AuthService struct {
	UserRepository *user.UserRepository
}

// NewAuthService создаёт сервис с репозиторием пользователей.
func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

// Register создаёт пользователя с хешем пароля и возвращает его email.
func (service *AuthService) Register(email, password, name string) (string, error) {

	// Проверяем, что никнейм свободен; отсутствие записи при регистрации ожидаемо.
	existedUser, err := service.UserRepository.GetByName(name)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", errors.New("запись не найдена")
	}
	if existedUser != nil {
		return "", errors.New(ErrUserNameExists)
	}

	// Проверяем, что email ещё не занят другим пользователем.
	existedUser, err = service.UserRepository.GetByEmail(email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", errors.New("запись не найдена")
	}
	if existedUser != nil {
		return "", errors.New(ErrUserEmailExists)
	}

	// Хешируем пароль с автоматически созданной солью перед сохранением в БД.
	cryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// В модель записываем только хеш, исходный пароль не сохраняем.
	user := &user.User{
		Email:    email,
		Password: string(cryptedPassword),
		Name:     name,
	}

	// Сохраняем пользователя; при ошибке не сообщаем об успешной регистрации.
	if _, err = service.UserRepository.Create(user); err != nil {
		return "", errors.New(ErrRegistrationErr)
	}
	return user.Email, nil
}

// Login находит пользователя по никнейму или email и проверяет его пароль.
func (service *AuthService) Login(email, password, name string) (string, error) {
	// Для входа нужны пароль и хотя бы один идентификатор: email или никнейм.
	if (email == "" && name == "") || password == "" {
		return "", errors.New(ErrNotEnougthData)
	}

	var existedUser *user.User
	var err error
	// Ищем по никнейму, если он указан; иначе используем email.
	if name != "" {
		existedUser, err = service.UserRepository.GetByName(name)
	} else {
		existedUser, err = service.UserRepository.GetByEmail(email)
	}

	// Отсутствие пользователя скрываем за общей ошибкой входа, ошибки БД сохраняем.
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New(ErrWrongCredetials)
		}
		return "", err
	}

	// Защищаемся от обращения к полям nil, даже если репозиторий не вернул ошибку.
	if existedUser == nil {
		return "", errors.New(ErrWrongCredetials)
	}

	// Сравниваем хеш из БД с исходным паролем; причина отказа остаётся общей.
	if err = bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(password)); err != nil {
		return "", errors.New(ErrWrongCredetials)
	}

	return existedUser.Email, nil
}
