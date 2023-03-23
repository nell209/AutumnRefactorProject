package services

import (
	"context"
	"github.com/nell209/AutumnRefactor/models"
)

func (s *Service) GetUserFromContext(ctx context.Context) (models.User, error) {
	var user models.User
	userID := ctx.Value(UserIDKey).(string)
	err := s.db.First(&user, "id = ?", userID).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
