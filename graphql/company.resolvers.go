package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/nell209/AutumnRefactor/models"

	"github.com/nell209/AutumnRefactor/graphql/generated"
)

// Branches is the resolver for the Branches field.
func (r *companyResolver) Branches(ctx context.Context, obj *models.Company) ([]*models.Branch, error) {
	panic(fmt.Errorf("not implemented"))
}

// Company returns generated.CompanyResolver implementation.
func (r *Resolver) Company() generated.CompanyResolver { return &companyResolver{r} }

type companyResolver struct{ *Resolver }
