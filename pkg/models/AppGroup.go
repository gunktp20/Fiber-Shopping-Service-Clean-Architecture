package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AppGroup struct {
	ID          string `gorm:"type:uuid;primaryKey;not null"`
	Name        string `gorm:"unique;not null" json:"name"`
	Description string `gorm:"not null" json:"description"`
	IconURL     string `gorm:"unique;not null" json:"icon_url"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Solutions   []Solution
}

func (appGroup *AppGroup) BeforeCreate(tx *gorm.DB) (err error) {
	if appGroup.ID == "" {
		appGroup.ID = uuid.New().String() // สร้าง UUID อัตโนมัติ
	}
	return
}
