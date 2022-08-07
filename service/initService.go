package service

import (
	"context"
	"github.com/nell209/AutumnRefactor/graph/model"
)

// Init create a mock company and branch for quick testing purposes
func (s *Service) Init(ctx context.Context) error {
	branch := model.Branch{
		Name: "My Branch",
	}
	company := model.Company{
		Name:     "My Company",
		Branches: []*model.Branch{&branch},
	}
	if err := s.db.Create(&company).Error; err != nil {
		// log the error here
		return err
	}
	return nil
}
