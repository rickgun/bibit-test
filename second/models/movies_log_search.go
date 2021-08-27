package models

import (
	"gorm.io/gorm"
)

type MoviesLogSearch struct {
	gorm.Model
	ID   int64  `gorm:"primaryKey"`
	Data string `gorm:"column:payload;type:json" json:"data"`
}

func (MoviesLogSearch) TableName() string {
	return "movies_log_search"
}
