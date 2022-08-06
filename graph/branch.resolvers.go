package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/nell209/AutumnRefactor/graph/generated"
	"github.com/nell209/AutumnRefactor/graph/model"
)

// TaskPriorities is the resolver for the taskPriorities field.
func (r *branchResolver) TaskPriorities(ctx context.Context, obj *model.Branch) ([]*model.TaskPriority, error) {
	panic(fmt.Errorf("not implemented"))
}

// Branch returns generated.BranchResolver implementation.
func (r *Resolver) Branch() generated.BranchResolver { return &branchResolver{r} }

type branchResolver struct{ *Resolver }