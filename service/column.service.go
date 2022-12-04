package service

import (
	"github.com/getsentry/sentry-go"
	"github.com/nell209/AutumnRefactor/graph/model"
)

func (s *Service) GetBranchColumns(branchId string) ([]*model.Column, error) {
	var columns []*model.Column
	if err := s.db.Where("branch_id = ?", branchId).Find(&columns).Error; err != nil {
		sentry.CaptureException(err)
		return columns, err
	}
	return columns, nil
}
