package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/poloping/interview_chenping_20240731/config"
	"github.com/poloping/interview_chenping_20240731/controller"
	docs "github.com/poloping/interview_chenping_20240731/docs"
	"github.com/poloping/interview_chenping_20240731/helper"
	"github.com/poloping/interview_chenping_20240731/model"
	"github.com/poloping/interview_chenping_20240731/repository"
	"github.com/poloping/interview_chenping_20240731/service"
	"github.com/rs/zerolog/log"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// @title  Service API
// @version	1.0
// @description A demo service API in Go using Gin framework

// @host 	localhost:8888
// @BasePath /api
func main() {

	log.Info().Msg("Started Server!")
	// Database
	db := config.DatabaseConnection()

	validate := validator.New()

	db.AutoMigrate(&model.Level{}, &model.Player{})

	// Repository
	playerRepository := repository.NewPlayerRepositoryImpl(db)
	levelRepository := repository.NewLevelRepositoryImpl(db)

	// Service
	playerService := service.NewPlayerServiceImpl(playerRepository, validate)
	levelService := service.NewLevelServiceImpl(levelRepository, validate)

	// Controller
	playerController := controller.NewPlayerController(playerService)
	levelController := controller.NewLevelController(levelService)

	// Router
	router := gin.Default()

	// add swagger
	docs.SwaggerInfo.BasePath = "/api/v1"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	baseRouter := router.Group("/api/v1")
	playerRouter := baseRouter.Group("/players")
	playerRouter.GET("", playerController.FindAll)
	playerRouter.GET("/:id", playerController.FindById)
	playerRouter.POST("", playerController.Create)
	playerRouter.PUT("/:id", playerController.Update)
	playerRouter.DELETE("/:id", playerController.Delete)

	levelRouter := baseRouter.Group("/levels")
	levelRouter.GET("", levelController.FindAll)
	levelRouter.POST("", levelController.Create)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
