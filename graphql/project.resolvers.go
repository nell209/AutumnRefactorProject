package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/nell209/AutumnRefactor/models"

	"github.com/nell209/AutumnRefactor/graphql/generated"
)

// Tasks is the resolver for the tasks field.
func (r *projectResolver) Tasks(ctx context.Context, obj *models.Project) ([]*models.Task, error) {
	tasks, err := r.service.GetProjectTasks(obj.ID)
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

// Project returns generated.ProjectResolver implementation.
func (r *Resolver) Project() generated.ProjectResolver { return &projectResolver{r} }

type projectResolver struct{ *Resolver }
