package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"github.com/nell209/AutumnRefactor/models"

	"github.com/nell209/AutumnRefactor/graphql/generated"
)

// Columns is the resolver for the columns field.
func (r *branchResolver) Columns(ctx context.Context, obj *models.Branch) ([]*models.Column, error) {
	columns, err := r.service.GetBranchColumns(obj.ID)
	if err != nil {
		return nil, errors.New("couldn't retrieve the columns")
	}
	return columns, nil
}

// Users is the resolver for the users field.
func (r *branchResolver) Users(ctx context.Context, obj *models.Branch) ([]*models.User, error) {
	users, err := r.service.GetBranchUsers(obj.ID)
	if err != nil {
		return users, err
	}
	return users, nil
}

// Projects is the resolver for the projects field.
func (r *branchResolver) Projects(ctx context.Context, obj *models.Branch) ([]*models.Project, error) {
	projects, err := r.service.GetBranchProjects(obj.ID)
	if err != nil {
		return projects, err
	}
	return projects, nil
}

// Branch returns generated.BranchResolver implementation.
func (r *Resolver) Branch() generated.BranchResolver { return &branchResolver{r} }

type branchResolver struct{ *Resolver }
