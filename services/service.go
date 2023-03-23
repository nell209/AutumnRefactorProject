package services

import "gorm.io/gorm"

// Service Database is injected at the services level to ensure it's a services level singleton
// and ensure separation of concerns on the api level
// I see no reason resolvers would need direct access to the DB
type Service struct {
	db *gorm.DB
}

func InitializeServices(db *gorm.DB) *Service {
	return &Service{db: db}
}
