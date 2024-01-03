package handler

import (
	"context"
	"errors"

	models "github.com/tylerneath/nuture-backend/internal/model"
	"gorm.io/gorm"
)

type (
	UserHandler struct {
		db gorm.DB
	}
	MessageHandler struct {
		db gorm.DB
	}
)

func New(ctx context.Context, db gorm.DB) (*UserHandler, *MessageHandler) {
	return &UserHandler{
			db: db,
		}, &MessageHandler{
			db: db,
		}
}

func (v *UserHandler) createUser() error {
	return errors.New("implement me")
}

func (v *UserHandler) deleteUser() error {
	return errors.New("implement me")
}

func (v *UserHandler) getUser() (models.User, error) {
	var user models.User
	return user, errors.New("implement me")
}

func (v *MessageHandler) createMessage() error {
	return errors.New("implement me")
}
