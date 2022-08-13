package service

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/nell209/AutumnRefactor/graph/model"
	"golang.org/x/crypto/bcrypt"
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
func (s *AuthService) SignAuthJWT(userID string, role string) (string, error) {
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

func (s *AuthService) ValidateJWT(signedToken string) (AutumnClaims, error) {
	var parsedClaims AutumnClaims
	token, err := jwt.ParseWithClaims(signedToken, &parsedClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	if !token.Valid {
		return parsedClaims, jwt.ErrTokenUnverifiable
	}
	return parsedClaims, err
}

// take a value and return an updated copy
func hashPassword(password string, user model.User) (model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, err
	}
	user.Password = string(hashedPassword)
	return user, nil
}

func checkPassword(providedPassword string, user model.User) error {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword)); err != nil {
		return err
	}
	return nil
}
