package restHandler

import "github.com/nell209/AutumnRefactor/services"

type restHandler struct {
	service *services.Service
}

func InitializeHandler(service *services.Service) restHandler {
	return restHandler{service: service}
}
