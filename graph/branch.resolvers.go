package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/nell209/AutumnRefactor/graph/generated"
	"github.com/nell209/AutumnRefactor/graph/model"
)

// Columns is the resolver for the columns field.
func (r *branchResolver) Columns(ctx context.Context, obj *model.Branch) ([]*model.Column, error) {
	panic(fmt.Errorf("not implemented"))
}

// Branch returns generated.BranchResolver implementation.
func (r *Resolver) Branch() generated.BranchResolver { return &branchResolver{r} }

type branchResolver struct{ *Resolver }
