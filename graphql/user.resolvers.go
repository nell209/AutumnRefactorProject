package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/nell209/AutumnRefactor/models"

	"github.com/nell209/AutumnRefactor/graphql/generated"
)

// Tasks is the resolver for the tasks field.
func (r *userResolver) Tasks(ctx context.Context, obj *models.User) ([]*models.Task, error) {
	tasks, err := r.service.GetUserTasks(obj.ID)
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
