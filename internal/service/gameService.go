package service

import (
	"github.com/d-chosen-one/tgb/internal/creator"
	"github.com/d-chosen-one/tgb/internal/db"
	"github.com/d-chosen-one/tgb/internal/model"
	"github.com/google/uuid"
)

type gameService struct {
	repository db.GameRepository
}

func New(repository db.GameRepository) Service {
	instance := new(gameService)
	instance.repository = repository
	return *instance
}

func (g gameService) GetAllGames() ([]*model.Game, error) {
	games, err := g.repository.GetAllGames()
	if err != nil {
		return nil, err
	}
	return games, nil

}

func (g gameService) CreateNewGame(gameSettings model.GameSettings) (*model.Game, error) {
	game := creator.NewGame(gameSettings.GameField)
	game.Id = uuid.New().String()
	game.PlayerName = gameSettings.PlayerName
	game.PlayerRace = gameSettings.PlayerRace
	err := g.repository.SaveGame(game)
	if err != nil {
		return nil, err
	}
	return &game, nil
}
