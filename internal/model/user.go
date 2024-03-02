package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Base
	Username       *string   `gorm:"uniqueIndex;index:idx_username;type:varchar(100)" json:"username"`
	Email          string    `gorm:"uniqueIndex;default:null;index:idx_email_null;type:varchar(100);not null" json:"email"`
	Messages       []Message `gorm:"type:uuid;OnDelete:CASCADE;foreignKey:UserID" json:"messages"`
	HashedPassword string    `json:"-"`
}

func DeleteUser(db *gorm.DB, username string) error {
	panic("implement me")
}

func GetUser(db *gorm.DB, userID uuid.UUID) (*User, error) {
	panic("impplement me")
}
