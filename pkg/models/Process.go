package models

import "time"

type Process struct {
	ID           string    `gorm:"type:uuid;primaryKey;not null"`
	Name         string    `gorm:"unique;not null"`
	Description  string    `gorm:"not null"`
	SolutionID   string    `gorm:"type:uuid;not null"`
	Solution     Solution  `gorm:"foreignKey:SolutionID"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
	SubProcesses []SubProcess
}
