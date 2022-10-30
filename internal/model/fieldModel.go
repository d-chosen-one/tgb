package model

import "github.com/d-chosen-one/tgb/internal/enums"

type Field struct {
	Position Position     `json:"position"`
	Owner    enums.ERaces `json:"owner"`
}

type Position struct {
	Row    int `json:"row"`
	Column int `json:"column"`
}
