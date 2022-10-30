package service

import (
	"github.com/d-chosen-one/tgb/internal/enums"
	"github.com/d-chosen-one/tgb/internal/model"
)

//func (s *gameService) CreateNewGame(gSettings model.GameSettings) model.Game {

//}
//
func createFields(gSettings model.GameSettings) []model.Field {
	fields := make([]model.Field, 0)
	switch gSettings.GameField {
	case enums.Small:
		fields = makeFields(10, 10)
	case enums.Medium:
		fields = makeFields(20, 20)
	case enums.Large:
		fields = makeFields(30, 30)
	default:
		fields = makeFields(5, 5)
	}
	return fields
}

func makeFields(rows, columns int) []model.Field {

	var fields = make([]model.Field, 0)
	for row := 1; row <= rows; row++ {
		for column := 1; column <= columns; column++ {
			var position = new(model.Position)
			position.Row = row
			position.Column = column
			var field = new(model.Field)
			field.Position = *position
			fields = append(fields, *field)
		}
	}
	return fields
}
