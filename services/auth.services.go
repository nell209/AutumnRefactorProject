package services

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nell209/AutumnRefactor/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

// TODO signed secret should be complex and held in environment variables
const tokenSecret string = "blah blah blah"

type AuthService struct {
	signedSecret        *string
	expiredTimeInSecond *time.Duration
}

type AutumnClaims struct {
	Role   string
	UserID string
	jwt.RegisteredClaims
}

// UserIDKey A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var UserIDKey = &contextKey{"UserID"}
var RoleKey = &contextKey{"Role"}

type contextKey struct {
	name string
}

func NewAuthService() AuthService {
	return AuthService{
		//&config.JWTSecret,
		//&config.JWTExpireIn
	}
}

// SignAuthJWT TODO decide an expire time
func SignAuthJWT(userID string, role string) (string, error) {
	expirationTime := jwt.NumericDate{Time: time.Now().Add(time.Second * 40000)}
	claims := AutumnClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &expirationTime,
		},
	}
	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := unsignedToken.SignedString([]byte(tokenSecret))
	return signedToken, err
}

//func (s *AuthService) SignRenewJWT(user *models.User) (*string, error) {
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
//		"userID": base64.StdEncoding.EncodeToString([]byte(user.Uid)),
//		"type":   base64.StdEncoding.EncodeToString([]byte(constants.JWTRefresh)),
//		"exp":    time.Now().AddDate(0, 6, 0).Unix(), // lasts 6 months
//	})
//	tokenString, err := token.SignedString([]byte(*s.signedSecret))
//	return &tokenString, err
//}

func ValidateJWT(signedToken string) (AutumnClaims, error) {
	var parsedClaims AutumnClaims
	token, err := jwt.ParseWithClaims(signedToken, &parsedClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	if token == nil || !token.Valid {
		return parsedClaims, jwt.ErrTokenUnverifiable
	}
	return parsedClaims, err
}

// TODO remove this before Launch
func (s *Service) AdminSetPassword(userID string, password string) error {
	var user models.User
	hashedPass, err := HashPassword(password)
	if err != nil {
		return err
	}
	err = s.db.Model(&user).Where("id = ?", userID).Update("password", hashedPass).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// take a value and return an updated copy
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func checkPassword(providedPassword string, dbPassword *string) error {
	if dbPassword == nil {
		return errors.New("no password set to compare against")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(*dbPassword), []byte(providedPassword)); err != nil {
		return err
	}
	return nil
}

func (s *Service) checkAndAuthenticateTempAccount(userID string, hashedPassword string) (models.TemporaryAccount, error) {
	tempAccount := models.TemporaryAccount{UserID: userID}
	err := s.db.Where(&tempAccount).Joins("User").First(&tempAccount).Error
	if err != nil {
		return tempAccount, err
	}
	err = checkPassword(hashedPassword, &tempAccount.TemporaryPassword)
	if err != nil {
		return tempAccount, err
	}
	return tempAccount, nil
}

func (s *Service) AuthenticateCredentials(email string, password string) (models.User, *string, error) {
	var user models.User
	if err := s.db.Where("LOWER(email) = LOWER(?)", email).First(&user).Error; err != nil {
		return user, user.Role, err
	}
	if user.Password == nil {
		tempAccount, err := s.checkAndAuthenticateTempAccount(user.ID, password)
		if err != nil {
			return user, user.Role, err
		}
		tempRole := models.TEMP
		return tempAccount.User, &tempRole, nil
	}
	if err := checkPassword(password, user.Password); err != nil {
		return user, user.Role, err
	}
	return user, user.Role, nil
}
