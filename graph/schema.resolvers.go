package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"github.com/nell209/AutumnRefactor/graph/generated"
	"github.com/nell209/AutumnRefactor/graph/model"
)

// Branch is the resolver for the branch field.
func (r *queryResolver) Branch(ctx context.Context, branchID string) (*model.Branch, error) {
	branch, err := r.service.GetBranch(branchID)
	if err != nil {
		return nil, errors.New("couldn't query for branch")
	}
	return &branch, nil
}

// GetUnassignedTasks is the resolver for the getUnassignedTasks field.
func (r *queryResolver) GetUnassignedTasks(ctx context.Context, branchID string) ([]*model.Task, error) {
	tasks, err := r.service.GetUnassignedTasks(branchID)
	if err != nil {
		return tasks, errors.New("could not fetch unassigned tasks")
	}
	return tasks, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
