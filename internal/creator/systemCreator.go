package creator

import (
	"fmt"
	"github.com/d-chosen-one/tgb/internal/enums"
	"github.com/d-chosen-one/tgb/internal/model"
	"math/rand"
)

func (gc *GameCreator) setSystems() {
	var position model.Position
	for index := 0; index < gc.SystemCount; index++ {
		do := true
		for do {
			row := rand.Intn(gc.RowCount) + 1
			column := rand.Intn(gc.RowCount) + 1
			fmt.Printf("row: %v column: %v \\n", row, column)

			position = model.NewPosition(row, column)
			if !gc.Fields[position].IsSystem() {
				gc.Fields[position].Kind = enums.System
				do = false
			}
		}
	}
}
