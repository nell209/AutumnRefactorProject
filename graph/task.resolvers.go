package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/nell209/AutumnRefactor/graph/generated"
	"github.com/nell209/AutumnRefactor/graph/model"
)

// Users is the resolver for the users field.
func (r *taskResolver) Users(ctx context.Context, obj *model.Task) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Prerequisites is the resolver for the prerequisites field.
func (r *taskResolver) Prerequisites(ctx context.Context, obj *model.Task) ([]*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

// Task returns generated.TaskResolver implementation.
func (r *Resolver) Task() generated.TaskResolver { return &taskResolver{r} }

type taskResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *taskResolver) Postrequisites(ctx context.Context, obj *model.Task) ([]*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *taskResolver) StartedBy(ctx context.Context, obj *model.Task) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}
