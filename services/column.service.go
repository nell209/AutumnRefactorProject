package services

import (
	"github.com/getsentry/sentry-go"
	"github.com/nell209/AutumnRefactor/models"
)

func (s *Service) GetBranchColumns(branchId string) ([]*models.Column, error) {
	var columns []*models.Column
	if err := s.db.Where("branch_id = ?", branchId).Find(&columns).Error; err != nil {
		sentry.CaptureException(err)
		return columns, err
	}
	return columns, nil
}
