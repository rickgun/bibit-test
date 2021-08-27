package repository

import (
	MoviesLogInterface "bibit-test/app/movies-log"
	"bibit-test/models"

	"github.com/jinzhu/gorm"
)

type MoviesLogRepository struct {
	Conn *gorm.DB
}

func NewMoviesLogRepository(Conn *gorm.DB) MoviesLogInterface.IMoviesLogRepository {
	return &MoviesLogRepository{Conn}
}

func (m *MoviesLogRepository) InsertSearch(data *models.MoviesLogSearch) error {
	tx := m.Conn.Begin()
	tx.Model(&data).Create(&data)
	tx.Commit()

	return tx.Error
}

func (m *MoviesLogRepository) InsertDetail(data *models.MoviesLogDetail) error {
	tx := m.Conn.Begin()
	tx.Model(&data).Create(&data)
	tx.Commit()

	return tx.Error
}
