package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/nell209/AutumnRefactor/graph/generated"
	"github.com/nell209/AutumnRefactor/graph/model"
)

// DeletedBy is the resolver for the deletedBy field.
func (r *projectResolver) DeletedBy(ctx context.Context, obj *model.Project) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Tasks is the resolver for the tasks field.
func (r *projectResolver) Tasks(ctx context.Context, obj *model.Project) ([]*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

// Users is the resolver for the users field.
func (r *projectResolver) Users(ctx context.Context, obj *model.Project) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Project returns generated.ProjectResolver implementation.
func (r *Resolver) Project() generated.ProjectResolver { return &projectResolver{r} }

type projectResolver struct{ *Resolver }
