package service

import "github.com/d-chosen-one/tgb/internal/model"

type Service interface {
	FindAll() []model.Game
	SaveOne(value model.Game) model.Game
}

type service struct {
	Values []model.Game
}

func (s *service) FindAll() []model.Game {
	return s.Values
}

func (s *service) SaveOne(value model.Game) model.Game {
	s.Values = append(s.Values, value)
	return value
}

func New() Service {
	return new(service)
}
