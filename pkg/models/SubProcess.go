package models

import "time"

type SubProcess struct {
	ID          string  `gorm:"type:uuid;primaryKey;not null"`
	Name        string  `gorm:"unique;not null"`
	Description string  `gorm:"not null"`
	ProcessID   string  `gorm:"type:uuid;not null"`
	Process     Process `gorm:"foreignKey:ProcessID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Types       []Type
}
