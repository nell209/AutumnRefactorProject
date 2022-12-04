package service

import (
	"github.com/getsentry/sentry-go"
	"github.com/nell209/AutumnRefactor/graph/model"
	"gorm.io/gorm/clause"
)

// todo, add context checking to these

func (s *Service) CreateProject(project model.Project) (model.Project, error) {
	err := s.db.Create(&project).Error
	if err != nil {
		return project, err
	}
	return project, nil
}

func (s *Service) UpdateProject(projectID string, project model.Project) (model.Project, error) {
	err := s.db.Model(&project).Clauses(clause.Returning{}).Where("id = ?", projectID).Updates(&project).Error
	if err != nil {
		return project, err
	}
	return project, nil
}

func (s *Service) GetTaskProject(projectId string) (model.Project, error) {
	var project model.Project
	if err := s.db.Where("id = ?", projectId).First(&project).Error; err != nil {
		sentry.CaptureException(err)
		return project, err
	}
	return project, nil
}
