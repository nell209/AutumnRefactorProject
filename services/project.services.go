package services

import (
	"github.com/getsentry/sentry-go"
	"github.com/nell209/AutumnRefactor/models"
	"gorm.io/gorm/clause"
)

// todo, add context checking to these

func (s *Service) CreateProject(project models.Project) (models.Project, error) {
	err := s.db.Create(&project).Error
	if err != nil {
		return project, err
	}
	return project, nil
}

func (s *Service) UpdateProject(projectID string, project models.Project) (models.Project, error) {
	err := s.db.Model(&project).Clauses(clause.Returning{}).Where("id = ?", projectID).Updates(&project).Error
	if err != nil {
		return project, err
	}
	return project, nil
}

func (s *Service) GetTaskProject(projectId string) (models.Project, error) {
	var project models.Project
	if err := s.db.Where("id = ?", projectId).First(&project).Error; err != nil {
		sentry.CaptureException(err)
		return project, err
	}
	return project, nil
}

func (s *Service) GetBranchProjects(branchId string) ([]*models.Project, error) {
	var projects []*models.Project
	err := s.db.Where("branch_id = ?", branchId).Error
	if err != nil {
		return projects, err
	}
	return projects, nil
}
