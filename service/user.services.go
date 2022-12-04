package service

import (
	"context"
	"errors"
	"github.com/nell209/AutumnRefactor/graph/model"
	"log"
)

type userInfo struct {
	firstName string
	lastName  string
	email     string
	password  string
}

func (s *Service) CreateUser(ctx context.Context, info userInfo) (model.User, error) {
	// check is ADMIN
	user := model.User{
		FirstName:        info.firstName,
		LastName:         info.lastName,
		Email:            info.email,
		Password:         &info.password,
		DefaultFilterUID: nil,
	}
	if err := s.db.Create(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (s *Service) CreateTempUser(ctx context.Context, email string, password *string) (model.User, error) {
	// TODO if no password, use the company password
	if password == nil {
		return model.User{}, errors.New("please provide a password for now")
	}
	hashedPassword, err := HashPassword(*password)
	if err != nil {
		log.Println()
		return model.User{}, errors.New("couldn't hash your password")
	}
	callerUserID := ctx.Value(UserIDKey).(string)
	caller, err := s.GetUserByID(callerUserID)
	if err != nil {
		return model.User{}, err
	}
	role := model.TEMP
	user := model.User{
		Email:    email,
		Role:     &role,
		BranchID: caller.BranchID,
	}
	account := model.TemporaryAccount{
		User:              user,
		TemporaryPassword: hashedPassword,
		Authenticated:     false,
	}
	if err := s.db.Create(&account).Error; err != nil {
		return account.User, err
	}
	return account.User, nil
}

func (s *Service) GetUserByID(userID string) (model.User, error) {
	var user model.User
	err := s.db.First(&user, "id = ?", userID).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// SetPermanentPassword TODO, could be optimized to be a single transaction
func (s *Service) SetPermanentPassword(userID string, password string, firstName string, lastName string) error {
	var tempAccount model.TemporaryAccount
	err := s.db.Model(&tempAccount).Where("user_id = ?", userID).Joins("User").First(&tempAccount).Error
	if tempAccount.Authenticated == true || tempAccount.User.Role == nil || *tempAccount.User.Role != model.TEMP {
		return errors.New("already Authenticated")
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}
	err = s.db.Model(&tempAccount.User).Where("id = ?", userID).Select("Password", "FirstName", "LastName").Updates(model.User{Password: &hashedPassword, FirstName: firstName, LastName: lastName}).Error
	if err != nil {
		log.Println(err)
		return err
	}
	if err := s.PromoteTempAccount(tempAccount); err != nil {
		return err
	}
	return nil
}

func (s *Service) PromoteTempAccount(tempAccount model.TemporaryAccount) error {
	err := s.db.Model(&tempAccount).Where("id = ?", tempAccount.ID).Update("authenticated", true).Error
	if err != nil {
		return err
	}
	err = s.db.Model(&tempAccount.User).Where("id = ?", tempAccount.UserID).Update("role", model.BASIC).Error
	if err != nil {
		return err
	}
	return nil
}
