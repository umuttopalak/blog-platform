package responses

import "github.com/google/uuid"

type UserResponse struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email,omitempty"` // Boş ise gösterilmez
}

type RegisterResponse struct {
	Message string       `json:"message"`
	User    UserResponse `json:"user"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
