package service

import (
	"errors"
	"fmt"
)

var (
	ErrUserNotFound        = errors.New("user not found")
	ErrUserAlreadyExists   = errors.New("user already exists")
	ErrInvalidPassword     = errors.New("invalid password")
	ErrInvalidRequest      = errors.New("invalid request")
	ErrInternalServerError = errors.New("internal server error")
	ErrInvalidToken        = errors.New("invalid token")
	ErrMessageNotFound     = errors.New("message not found")
	ErrInvalidID           = errors.New("invalid id")
	ErrMissingField        = errors.New("missing field")
)

func NewInternalServerError(err error) error {
	return fmt.Errorf("internal server error: %w", err)
}
