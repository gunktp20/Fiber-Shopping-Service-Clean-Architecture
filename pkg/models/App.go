package models

import "time"

type App struct {
	ID                 string   `gorm:"type:uuid;primaryKey;not null"`
	AppName            string   `gorm:"unique;not null"`
	AppLink            string   `gorm:"not null"`
	ImageURL           string   `gorm:"not null"`
	ShortDescription   string   `gorm:"not null"`
	KeyFeatures        []string `gorm:"type:text[]"`
	Overview           string
	AppGroupID         string   `gorm:"type:uuid;not null"`
	AppGroup           AppGroup `gorm:"foreignKey:AppGroupID"`
	SolutionID         string   `gorm:"type:uuid;not null"`
	Solution           Solution `gorm:"foreignKey:SolutionID"`
	TypeID             string   `gorm:"type:uuid;not null"`
	Type               Type     `gorm:"foreignKey:TypeID"`
	CostReduction      bool     `gorm:"default:false;not null"`
	EfficiencyIncrease bool     `gorm:"default:false;not null"`
	ProductionIncrease bool     `gorm:"default:false;not null"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
