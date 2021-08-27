package handler

import (
	"strconv"
	"time"

	Base "bibit-test/app/api/handler"
	MoviesInterface "bibit-test/app/movies"

	"github.com/gin-gonic/gin"
)

type MoviesResponse struct {
	HealthStatus string    `json:"health_status"`
	DBTimestamp  time.Time `json:"database_timestamp"`
}

type MoviesHandler struct {
	MoviesUsecase MoviesInterface.IMoviesUsecase
}

func (a *MoviesHandler) Search(c *gin.Context) {
	var (
		query = c.Request.URL.Query()
	)

	if query.Get("pagination") == "" && query.Get("searchword") == "" {
		Base.RespondFailValidation(c, "Require pagination and searchword")
		return
	}

	pagination, err := strconv.ParseInt(query.Get("pagination"), 10, 64)
	if err != nil {
		Base.RespondFailValidation(c, "Require pagination and searchword")
		return
	}

	searchword := query.Get("searchword")
	res := a.MoviesUsecase.Search(pagination, searchword)

	Base.RespondJSON(c, res)
	return
}

func (a *MoviesHandler) Detail(c *gin.Context) {
	id, found := c.Params.Get("id")
	if found != true {
		Base.RespondFailValidation(c, "Require imbd id ")
		return
	}

	res := a.MoviesUsecase.Detail(id)

	Base.RespondJSON(c, res)
	return
}
