package service

import "gorm.io/gorm"

type Service struct {
	db *gorm.DB
}

func InitializeServices(db *gorm.DB) *Service {
	return &Service{db: db}
}
