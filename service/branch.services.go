package service

import (
	"github.com/getsentry/sentry-go"
	"github.com/nell209/AutumnRefactor/graph/model"
)

func (s *Service) GetBranch(branchId string) (model.Branch, error) {
	var branch model.Branch
	err := s.db.Where("id = ?", branchId).First(&branch).Error
	if err != nil {
		sentry.CaptureException(err)
		return branch, err
	}
	return branch, nil
}
