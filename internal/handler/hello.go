package handler

import (
	"github.com/d-chosen-one/tgb/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SaveHandler(c *gin.Context) {
	var game = new(model.Game)
	c.BindJSON(game)
	bimbam.SaveOne(*game)
	c.IndentedJSON(http.StatusOK, game)
}

func GamesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, bimbam.FindAll())
}
