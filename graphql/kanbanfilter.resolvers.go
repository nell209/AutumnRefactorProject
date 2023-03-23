package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/nell209/AutumnRefactor/models"

	"github.com/nell209/AutumnRefactor/graphql/generated"
)

// Columns is the resolver for the columns field.
func (r *kanbanFilterResolver) Columns(ctx context.Context, obj *models.KanbanFilter) ([]*models.Column, error) {
	panic(fmt.Errorf("not implemented"))
}

// KanbanFilter returns generated.KanbanFilterResolver implementation.
func (r *Resolver) KanbanFilter() generated.KanbanFilterResolver { return &kanbanFilterResolver{r} }

type kanbanFilterResolver struct{ *Resolver }
