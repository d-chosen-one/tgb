package handler

import "github.com/d-chosen-one/tgb/internal/service"

var bimbam service.Service

func Init(service service.Service) {
	bimbam = service
}
