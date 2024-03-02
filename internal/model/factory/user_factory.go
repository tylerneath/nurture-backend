package factory

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	models "github.com/tylerneath/nuture-backend/internal/model"
)

func MakeFakeUser() (*models.User, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	username := gofakeit.Username()
	u := models.User{
		Base: models.Base{
			ID:        id,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Username: &username,
		Email:    gofakeit.Email(),
	}

	return &u, nil
}
