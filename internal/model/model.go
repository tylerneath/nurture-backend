package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        uuid.UUID      `gorm:"column:id;type:uuid;primary_key;" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"updated_at" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"deleted_at,index" json:"deletedAt"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return errors.New("unable to create uuid for object")
	}
	base.ID = uuid

	base.CreatedAt = time.Now()
	base.UpdatedAt = time.Now()
	return nil

}

func (base *Base) AfterUpdate(tx *gorm.DB) error {
	base.UpdatedAt = time.Now()
	return nil

}
