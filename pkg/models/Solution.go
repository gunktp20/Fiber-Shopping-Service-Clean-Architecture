package models

import "time"

type Solution struct {
	ID          string   `gorm:"type:uuid;primaryKey;not null"`
	Name        string   `gorm:"unique;not null"`
	Description string   `gorm:"not null"`
	AppGroupID  string   `gorm:"type:uuid;not null"`
	AppGroup    AppGroup `gorm:"foreignKey:AppGroupID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Processes   []Process
}
