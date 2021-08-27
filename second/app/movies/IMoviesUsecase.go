package movies

import "bibit-test/models"

type IMoviesUsecase interface {
	Search(pagination int64, searchword string) *models.MoviesSearchResponse
	Detail(id string) *models.MovieDetailResponse
}
