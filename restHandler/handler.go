package restHandler

import "github.com/nell209/AutumnRefactor/service"

type restHandler struct {
	service *service.Service
}

func InitializeHandler(service *service.Service) restHandler {
	return restHandler{service: service}
}
