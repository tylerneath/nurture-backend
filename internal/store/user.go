package store

import (
	models "github.com/tylerneath/nuture-backend/internal/model"
	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (s *UserStore) Create(u *models.User) error {
	return s.db.Create(u).Error
}

func (s *UserStore) Get(id string) (*models.User, error) {
	var user models.User
	err := s.db.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (s *UserStore) Update(u *models.User) error {
	return s.db.Save(u).Error
}

func (s *UserStore) Delete(id string) error {
	return s.db.Delete(&models.User{}, id).Error
}
