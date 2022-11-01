package model

type Position struct {
	Row    int `json:"row"`
	Column int `json:"column"`
}

func NewPosition(row, column int) Position {
	position := new(Position)
	position.Row = row
	position.Column = column
	return *position
}
