package repo

import (
	"github.com/google/uuid"
	models "github.com/tylerneath/nuture-backend/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(*models.User) error
	GetByID(uuid.UUID) (*models.User, error)
	GetByEmail(string) (*models.User, error)
}

type userRepository struct {
	log *zap.Logger
	db  *gorm.DB
}

func NewUserRepository(db *gorm.DB, log *zap.Logger) UserRepository {
	return &userRepository{
		db:  db,
		log: log,
	}
}

func (u *userRepository) Create(user *models.User) error {
	if err := u.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) GetByEmail(email string) (*models.User, error) {
	user := models.User{}
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) GetByID(id uuid.UUID) (*models.User, error) {
	user := models.User{}
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
