package creator

import (
	"github.com/d-chosen-one/tgb/internal/model"
)

func (gc *GameCreator) createFields() {
	gc.Fields = make(map[model.Position]*model.Field)
	for row := 1; row <= gc.RowCount; row++ {
		for column := 1; column <= gc.ColumnCount; column++ {
			position := model.NewPosition(row, column)
			var field = new(model.Field)
			field.Position = position
			gc.Fields[field.Position] = field
		}
	}
}
