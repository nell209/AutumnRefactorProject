package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/nell209/AutumnRefactor/graph/generated"
	"github.com/nell209/AutumnRefactor/graph/model"
)

// Columns is the resolver for the columns field.
func (r *branchResolver) Columns(ctx context.Context, obj *model.Branch) ([]*model.Column, error) {
	columns, err := r.service.GetBranchColumns(obj.ID)
	if err != nil {
		return nil, errors.New("couldn't retrieve the columns")
	}
	return columns, nil
}

// Users is the resolver for the users field.
func (r *branchResolver) Users(ctx context.Context, obj *model.Branch) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Projects is the resolver for the projects field.
func (r *branchResolver) Projects(ctx context.Context, obj *model.Branch) ([]*model.Project, error) {
	panic(fmt.Errorf("not implemented"))
}

// Branch returns generated.BranchResolver implementation.
func (r *Resolver) Branch() generated.BranchResolver { return &branchResolver{r} }

type branchResolver struct{ *Resolver }
