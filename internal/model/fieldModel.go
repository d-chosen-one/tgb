package model

import "github.com/d-chosen-one/tgb/internal/enums"

type Field struct {
	Position Position         `json:"position"`
	Kind     enums.EFieldKind `json:"kind"`
}

func (f *Field) IsSystem() bool {
	return f.Kind == enums.System
}
