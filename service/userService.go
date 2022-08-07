package service

import (
	"context"
	"github.com/nell209/AutumnRefactor/graph/model"
)

type userInfo struct {
	firstName string
	lastName  string
	email     string
	password  string
}

func (s *Service) CreateUser(ctx context.Context, info userInfo) (model.User, error) {
	// is ADMIN
	user := model.User{
		FirstName:        info.firstName,
		LastName:         info.lastName,
		Email:            info.email,
		Password:         info.password,
		DefaultFilterUID: nil,
	}
	if err := s.db.Create(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}
