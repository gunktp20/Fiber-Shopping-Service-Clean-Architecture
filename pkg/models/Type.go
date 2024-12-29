package models

import "time"

type Type struct {
	ID           string     `gorm:"type:uuid;primaryKey;not null"`
	Name         string     `gorm:"unique;not null"`
	Description  string     `gorm:"not null"`
	SubProcessID string     `gorm:"type:uuid;not null"`
	SubProcess   SubProcess `gorm:"foreignKey:SubProcessID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Apps         []App
}
