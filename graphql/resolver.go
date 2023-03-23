package graphql

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	"github.com/nell209/AutumnRefactor/services"
)

type Resolver struct {
	service *services.Service
}

func BindServicesToResolver(services *services.Service) *Resolver {
	return &Resolver{service: services}
}
