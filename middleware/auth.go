package middleware

import (
	"errors"
	"github.com/nell209/AutumnRefactor/services"
	"net/http"
	"strings"
)

// AuthMiddleware Role is added to context to minimize the number of DB calls
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//token, err := validateAuthHeader(r, services.NewAuthService())
		//jwtClaims, err := services.ValidateJWT(token)
		//if err != nil {
		//	http.Error(w, "Forbidden", http.StatusForbidden)
		//	log.Println(err)
		//	return
		//}
		ctx := r.Context()
		//ctx = context.WithValue(r.Context(), services.UserIDKey, jwtClaims.UserID)
		//ctx = context.WithValue(ctx, services.RoleKey, jwtClaims.Role)
		//log.Println(ctx)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func TemporaryPasswordAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//	TODO
		next.ServeHTTP(w, r)
	})
}

// TODO re-write this https://stackoverflow.com/questions/39518237/how-to-extract-and-verify-token-sent-from-frontend
func validateAuthHeader(r *http.Request, authService services.AuthService) (string, error) {
	var tokenString string
	keys, ok := r.URL.Query()["at"]
	if !ok || len(keys) < 1 {
		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		if len(auth) != 2 || auth[0] != "Bearer" {
			return tokenString, errors.New("CredentialsError")
		}
		tokenString = auth[1]
	} else {
		tokenString = keys[0]
	}
	return tokenString, nil
}
