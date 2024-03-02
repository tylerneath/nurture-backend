package common

import "github.com/google/uuid"

type RegisterUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateMessageRequest struct {
	UserID  uuid.UUID `json:"user_id"`
	Message string    `json:"message"`
}
