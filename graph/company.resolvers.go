package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/nell209/AutumnRefactor/graph/generated"
	"github.com/nell209/AutumnRefactor/graph/model"
)

// Branches is the resolver for the Branches field.
func (r *companyResolver) Branches(ctx context.Context, obj *model.Company) ([]*model.Branch, error) {
	panic(fmt.Errorf("not implemented"))
}

// Company returns generated.CompanyResolver implementation.
func (r *Resolver) Company() generated.CompanyResolver { return &companyResolver{r} }

type companyResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *companyResolver) Users(ctx context.Context, obj *model.Company) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}
