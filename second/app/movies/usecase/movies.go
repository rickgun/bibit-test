package usecase

import (
	MoviesInterface "bibit-test/app/movies"
	MoviesLogInterface "bibit-test/app/movies-log"
	"bibit-test/models"
	"encoding/json"
	"fmt"
)

type MoviesUsecase struct {
	MoviesRepository    MoviesInterface.IMoviesRepository
	MoviesLogRepository MoviesLogInterface.IMoviesLogRepository
}

func NewMoviesUsecase(m MoviesInterface.IMoviesRepository, ml MoviesLogInterface.IMoviesLogRepository) MoviesInterface.IMoviesUsecase {
	return &MoviesUsecase{
		MoviesRepository:    m,
		MoviesLogRepository: ml,
	}
}

func (a *MoviesUsecase) Search(pagination int64, searchword string) *models.MoviesSearchResponse {
	movies, err := a.MoviesRepository.Search(pagination, searchword)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	body, err := json.Marshal(movies)
	if err != nil {
		return nil
	}

	err = a.MoviesLogRepository.InsertSearch(&models.MoviesLogSearch{
		Data: string(body),
	})
	if err != nil {
		return nil
	}

	return movies
}

func (a *MoviesUsecase) Detail(id string) *models.MovieDetailResponse {
	movies, err := a.MoviesRepository.Detail(id)
	if err != nil {
		return nil
	}

	body, err := json.Marshal(movies)
	if err != nil {
		return nil
	}

	err = a.MoviesLogRepository.InsertDetail(&models.MoviesLogDetail{
		Data: string(body),
	})
	if err != nil {
		return nil
	}

	return movies
}
