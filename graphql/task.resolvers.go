package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"github.com/nell209/AutumnRefactor/models"

	"github.com/nell209/AutumnRefactor/graphql/generated"
)

// Project is the resolver for the project field.
func (r *taskResolver) Project(ctx context.Context, obj *models.Task) (*models.Project, error) {
	// TODO this should definitely have a dataloader
	if obj.ProjectID == nil {
		return nil, errors.New("no available project set")
	}
	project, err := r.service.GetTaskProject(*obj.ProjectID)
	if err != nil {
		return nil, errors.New("could not fetch project for task")
	}
	return &project, nil
}

// Users is the resolver for the users field.
func (r *taskResolver) Users(ctx context.Context, obj *models.Task) ([]*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Prerequisites is the resolver for the prerequisites field.
func (r *taskResolver) Prerequisites(ctx context.Context, obj *models.Task) ([]*models.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

// Task returns generated.TaskResolver implementation.
func (r *Resolver) Task() generated.TaskResolver { return &taskResolver{r} }

type taskResolver struct{ *Resolver }
