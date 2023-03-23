package services

import (
	"github.com/getsentry/sentry-go"
	"github.com/nell209/AutumnRefactor/models"
)

func (s *Service) CreateTask(task models.Task) (models.Task, error) {
	err := s.db.Model(&models.Task{}).Create(&task).Error
	if err != nil {
		return task, err
	}
	return task, nil
}

func (s *Service) GetColumnTasks(columnId string) ([]*models.Task, error) {
	var tasks []*models.Task
	if err := s.db.Where("column_id = ?", columnId).Find(&tasks).Error; err != nil {
		sentry.CaptureException(err)
		return tasks, err
	}
	return tasks, nil
}

func (s *Service) GetUnassignedTasks(branchId string) ([]*models.Task, error) {
	var tasks []*models.Task
	if err := s.db.Where("column_id IS NULL").Find(&tasks).Error; err != nil {
		sentry.CaptureException(err)
		return tasks, err
	}
	return tasks, nil
}

func (s *Service) GetProjectTasks(projectId string) ([]*models.Task, error) {
	var tasks []*models.Task
	if err := s.db.Where("project_id = ?", projectId).Find(&tasks).Error; err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (s *Service) GetUserTasks(userId string) ([]*models.Task, error) {
	var tasks []*models.Task
	err := s.db.Model(&models.User{}).Where("user_id = ? ", userId).Association("Tasks").Find(&tasks)
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}
