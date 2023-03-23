package services

import (
	"context"
	"errors"
	"github.com/nell209/AutumnRefactor/models"
	"log"
)

type userInfo struct {
	firstName string
	lastName  string
	email     string
	password  string
}

func (s *Service) CreateUser(ctx context.Context, info userInfo) (models.User, error) {
	// check is ADMIN
	user := models.User{
		FirstName: info.firstName,
		LastName:  info.lastName,
		Email:     info.email,
		Password:  &info.password,
	}
	if err := s.db.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *Service) CreateTempUser(ctx context.Context, email string, password *string) (models.User, error) {
	// TODO if no password, use the company password
	if password == nil {
		return models.User{}, errors.New("please provide a password for now")
	}
	hashedPassword, err := HashPassword(*password)
	if err != nil {
		log.Println()
		return models.User{}, errors.New("couldn't hash your password")
	}
	callerUserID := ctx.Value(UserIDKey).(string)
	caller, err := s.GetUserByID(callerUserID)
	if err != nil {
		return models.User{}, err
	}
	role := models.TEMP
	user := models.User{
		Email:    email,
		Role:     &role,
		BranchID: caller.BranchID,
	}
	account := models.TemporaryAccount{
		User:              user,
		TemporaryPassword: hashedPassword,
		Authenticated:     false,
	}
	if err := s.db.Create(&account).Error; err != nil {
		return account.User, err
	}
	return account.User, nil
}

func (s *Service) GetUserByID(userID string) (models.User, error) {
	var user models.User
	err := s.db.First(&user, "id = ?", userID).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// SetPermanentPassword TODO, could be optimized to be a single transaction
func (s *Service) SetPermanentPassword(userID string, password string, firstName string, lastName string) error {
	var tempAccount models.TemporaryAccount
	err := s.db.Model(&tempAccount).Where("user_id = ?", userID).Joins("User").First(&tempAccount).Error
	if tempAccount.Authenticated == true || tempAccount.User.Role == nil || *tempAccount.User.Role != models.TEMP {
		return errors.New("already Authenticated")
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}
	err = s.db.Model(&tempAccount.User).Where("id = ?", userID).Select("Password", "FirstName", "LastName").Updates(models.User{Password: &hashedPassword, FirstName: firstName, LastName: lastName}).Error
	if err != nil {
		log.Println(err)
		return err
	}
	if err := s.PromoteTempAccount(tempAccount); err != nil {
		return err
	}
	return nil
}

func (s *Service) PromoteTempAccount(tempAccount models.TemporaryAccount) error {
	err := s.db.Model(&tempAccount).Where("id = ?", tempAccount.ID).Update("authenticated", true).Error
	if err != nil {
		return err
	}
	err = s.db.Model(&tempAccount.User).Where("id = ?", tempAccount.UserID).Update("role", models.BASIC).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetBranchUsers(branchId string) ([]*models.User, error) {
	var users []*models.User
	err := s.db.Where("branch_id = ?", branchId).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
