package movieslog

import "bibit-test/models"

type IMoviesLogRepository interface {
	InsertSearch(data *models.MoviesLogSearch) error
	InsertDetail(data *models.MoviesLogDetail) error
}
