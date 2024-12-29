package repository

import (
	"gorm.io/gorm"
)

type appGroupGormRepository struct {
	db *gorm.DB
}

func NewAppGroupGormRepository(db *gorm.DB) AppGroupRepositoryService {
	return &appGroupGormRepository{db}
}
