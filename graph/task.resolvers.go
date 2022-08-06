package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/nell209/AutumnRefactor/graph/generated"
	"github.com/nell209/AutumnRefactor/graph/model"
)

// StartedBy is the resolver for the startedBy field.
func (r *taskResolver) StartedBy(ctx context.Context, obj *model.Task) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Users is the resolver for the users field.
func (r *taskResolver) Users(ctx context.Context, obj *model.Task) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Prerequisites is the resolver for the prerequisites field.
func (r *taskResolver) Prerequisites(ctx context.Context, obj *model.Task) ([]*model.TaskPrerequisite, error) {
	panic(fmt.Errorf("not implemented"))
}

// Postrequisites is the resolver for the postrequisites field.
func (r *taskResolver) Postrequisites(ctx context.Context, obj *model.Task) ([]*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

// Task returns generated.TaskResolver implementation.
func (r *Resolver) Task() generated.TaskResolver { return &taskResolver{r} }

type taskResolver struct{ *Resolver }
