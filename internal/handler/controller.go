package handler

import (
	"github.com/d-chosen-one/tgb/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SaveHandler(c *gin.Context) {

}

func GamesHandler(c *gin.Context) {
	games, err := gameService.GetAllGames()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, games)
}

func NewGame(c *gin.Context) {
	var gameSettings = new(model.GameSettings)
	if err := c.BindJSON(gameSettings); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
	game, err := gameService.CreateNewGame(*gameSettings)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusOK, game)
}
