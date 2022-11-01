package creator

import (
	"github.com/d-chosen-one/tgb/internal/enums"
	"github.com/d-chosen-one/tgb/internal/model"
)

type GameCreator struct {
	Game        model.Game
	Fields      map[model.Position]*model.Field
	RowCount    int
	ColumnCount int
	SystemCount int
}

func NewGame(fieldSize enums.EFieldSize) model.Game {
	game := new(model.Game)
	newGame := new(GameCreator)
	newGame.setCounts(fieldSize)
	newGame.createFields()
	newGame.setSystems()
	return *game
}

func (gc *GameCreator) setCounts(size enums.EFieldSize) (rowCount, columnCount, systemCount int) {
	switch size {
	case enums.Small:
		gc.RowCount = 5
		gc.ColumnCount = 5
		gc.SystemCount = 5
	case enums.Medium:
		gc.RowCount = 8
		gc.ColumnCount = 8
		gc.SystemCount = 25
	case enums.Large:
		gc.RowCount = 12
		gc.ColumnCount = 12
		gc.SystemCount = 50
	default:
		gc.RowCount = 8
		gc.ColumnCount = 8
		gc.SystemCount = 25
	}
	return rowCount, columnCount, systemCount
}
