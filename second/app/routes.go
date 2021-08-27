package app

import (
	"github.com/gin-gonic/gin"

	MoviesInterface "bibit-test/app/movies"

	MHandler "bibit-test/app/movies/handler"
)

// Define your route here
// Register the route on main.go with usecase as the parameter

func MoviesHttpHandler(r *gin.Engine, us MoviesInterface.IMoviesUsecase) {
	mHandler := &MHandler.MoviesHandler{
		MoviesUsecase: us,
	}

	route := r.Group("/bibit")
	route.GET("/movies", mHandler.Search)
	route.GET("/movie/:id", mHandler.Detail)
}
