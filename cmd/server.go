package main

import (
	"github.com/d-chosen-one/tgb/internal/db"
	"github.com/d-chosen-one/tgb/internal/handler"
	"github.com/d-chosen-one/tgb/internal/service"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	err := db.InitMongo("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	repository := db.New()
	gameService := service.New(repository)
	handler.Init(gameService)
	router := gin.Default()
	router.POST("/new", handler.NewGame)
	//router.POST("/new", handler.SaveHandler)
	router.GET("/games", handler.GamesHandler)

	router.Run("localhost:8080")
}
