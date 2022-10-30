package service

import "github.com/d-chosen-one/tgb/internal/model"

type Service interface {
	GetAllGames() ([]*model.Game, error)
	CreateNewGame(gameSettings model.GameSettings) (*model.Game, error)
}
