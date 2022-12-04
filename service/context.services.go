package service

import (
	"context"
	"github.com/nell209/AutumnRefactor/graph/model"
)

func (s *Service) GetUserFromContext(ctx context.Context) (model.User, error) {
	var user model.User
	userID := ctx.Value(UserIDKey).(string)
	err := s.db.First(&user, "id = ?", userID).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
