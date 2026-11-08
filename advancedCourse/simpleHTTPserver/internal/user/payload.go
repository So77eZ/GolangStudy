package user

// UserCreateRequest содержит данные для регистрации пользователя.
type UserCreateRequest struct {
	Email    string `json:"email" gorm:"index,email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// UserUpdateRequest содержит данные для обновления пользователя.
type UserUpdateRequest struct {
	Email    string `json:"email" gorm:"index,email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
