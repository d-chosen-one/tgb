package handler

import "github.com/d-chosen-one/tgb/internal/service"

var gameService service.Service

func Init(service service.Service) {
	gameService = service
}
