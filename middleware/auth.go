package middleware

import (
	"context"
	"errors"
	"github.com/nell209/AutumnRefactor/service"
	"net/http"
	"strings"
)

// AuthMiddleware Role is added to context to minimize the number of DB calls
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwtClaims, err := validateAuthHeader(r, service.NewAuthService())
		if err != nil {
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(r.Context(), service.UserIDKey, jwtClaims.UserID)
		ctx = context.WithValue(ctx, service.RoleKey, jwtClaims.Role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// TODO re-write this https://stackoverflow.com/questions/39518237/how-to-extract-and-verify-token-sent-from-frontend
func validateAuthHeader(r *http.Request, authService service.AuthService) (service.AutumnClaims, error) {
	var tokenString string
	keys, ok := r.URL.Query()["at"]
	if !ok || len(keys) < 1 {
		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		if len(auth) != 2 || auth[0] != "Bearer" {
			return service.AutumnClaims{}, errors.New("CredentialsError")
		}
		tokenString = auth[1]
	} else {
		tokenString = keys[0]
	}
	jwtClaims, err := authService.ValidateJWT(tokenString)
	return jwtClaims, err
}
