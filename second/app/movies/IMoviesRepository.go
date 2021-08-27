package movies

import "bibit-test/models"

type IMoviesRepository interface {
	Search(pagination int64, searchword string) (movie *models.MoviesSearchResponse, err error)
	Detail(id string) (movie *models.MovieDetailResponse, err error)
}
