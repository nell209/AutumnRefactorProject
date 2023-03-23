package services

import (
	"github.com/getsentry/sentry-go"
	"github.com/nell209/AutumnRefactor/models"
)

func (s *Service) GetBranch(branchId string) (models.Branch, error) {
	var branch models.Branch
	err := s.db.Where("id = ?", branchId).First(&branch).Error
	if err != nil {
		sentry.CaptureException(err)
		return branch, err
	}
	return branch, nil
}
