package db

import "github.com/d-chosen-one/tgb/internal/model"

type GameRepository interface {
	GetAllGames() ([]*model.Game, error)
	SaveGame(game model.Game) error
	LoadGame(id string) (*model.Game, error)
}
