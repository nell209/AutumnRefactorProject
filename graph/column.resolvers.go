package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/nell209/AutumnRefactor/graph/generated"
	"github.com/nell209/AutumnRefactor/graph/model"
)

// Tasks is the resolver for the Tasks field.
func (r *columnResolver) Tasks(ctx context.Context, obj *model.Column) ([]*model.Task, error) {
	tasks, err := r.service.GetColumnTasks(obj.ID)
	if err != nil {
		return nil, errors.New("could not Fetch Column Tasks")
	}
	return tasks, err
}

// Column returns generated.ColumnResolver implementation.
func (r *Resolver) Column() generated.ColumnResolver { return &columnResolver{r} }

type columnResolver struct{ *Resolver }
