package main

import (
	"fmt"
	"log"

	"bibit-test/config"
	gorm "bibit-test/db"

	MLRepository "bibit-test/app/movies-log/repository"
	MRepository "bibit-test/app/movies/repository"
	MUsecase "bibit-test/app/movies/usecase"

	"github.com/gin-gonic/gin"

	routes "bibit-test/app"
)

var appConfig = config.Config.App

func main() {
	r := gin.New()

	db := gorm.MysqlConn()

	option := &MRepository.Option{
		BaseURL: appConfig.Address,
		APIKey:  appConfig.APIKey,
	}

	// Repositories
	mr := MRepository.NewMoviesRepository(option)
	mlr := MLRepository.NewMoviesLogRepository(db)

	// Usecases
	mu := MUsecase.NewMoviesUsecase(mr, mlr)

	// Handlers
	routes.MoviesHttpHandler(r, mu)

	if err := r.Run(fmt.Sprintf(":%s", appConfig.HTTPPort)); err != nil {
		log.Fatal(err)
	}
}
