package auth

// LoginRequest структура запроса логина
type LoginRequest struct {
	Email    string `json:"email" validate:"required_without=Name,omitempty,email"`
	Name     string `json:"name" validate:"required_without=Email"`
	Password string `json:"password" validate:"required"`
}
type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterResponse struct {
	Token string `json:"token"`
	ID    uint   `json:"id"`
	// Email string `json:"email"`
	// Name  string `json:"name"`
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}
