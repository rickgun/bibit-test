package models

import (
	"gorm.io/gorm"
)

type MoviesLogDetail struct {
	gorm.Model
	ID   int64  `gorm:"primaryKey;AUTO_INCREMENT;column:id" json:"id"`
	Data string `gorm:"column:payload;type:json" json:"data"`
}

func (MoviesLogDetail) TableName() string {
	return "movies_log_detail"
}
