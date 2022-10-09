package main

import (
	"github.com/d-chosen-one/tgb/internal/handler"
	"github.com/d-chosen-one/tgb/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	gameService := service.New()
	handler.Init(gameService)
	router := gin.Default()
	router.POST("/new", handler.SaveHandler)
	router.GET("/games", handler.GamesHandler)
	router.Run("localhost:8080")
}
