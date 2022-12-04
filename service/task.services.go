package service

import (
	"github.com/getsentry/sentry-go"
	"github.com/nell209/AutumnRefactor/graph/model"
)

func (s *Service) CreateTask(task model.Task) (model.Task, error) {
	err := s.db.Model(&model.Task{}).Create(&task).Error
	if err != nil {
		return task, err
	}
	return task, nil
}

func (s *Service) GetColumnTasks(columnId string) ([]*model.Task, error) {
	var tasks []*model.Task
	if err := s.db.Where("column_id = ?", columnId).Find(&tasks).Error; err != nil {
		sentry.CaptureException(err)
		return tasks, err
	}
	return tasks, nil
}

func (s *Service) GetUnassignedTasks(branchId string) ([]*model.Task, error) {
	var tasks []*model.Task
	if err := s.db.Where("column_id IS NULL").Find(&tasks).Error; err != nil {
		sentry.CaptureException(err)
		return tasks, err
	}
	return tasks, nil
}
