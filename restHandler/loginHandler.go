package restHandler

import (
	"encoding/json"
	"github.com/nell209/AutumnRefactor/service"
	"log"
	"net/http"
)

type loginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	AccessToken string `json:"accessToken"`
}

type t struct {
	UserID   string `json:"userID"`
	Password string `json:"password"`
}

func (h *restHandler) SetPassword() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		decoder := json.NewDecoder(r.Body)
		var body t
		err := decoder.Decode(&body)
		if err != nil {
			panic(err)
		}

		if err := h.service.AdminSetPassword(body.UserID, body.Password); err != nil {
			log.Println(err)
			http.Error(w, "couldn't set password", http.StatusInternalServerError)
		}
		w.Write([]byte("success"))
	})
}

func (h *restHandler) Login() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		decoder := json.NewDecoder(r.Body)
		var body loginCredentials
		err := decoder.Decode(&body)
		if err != nil {
			panic(err)
		}

		user, role, err := h.service.AuthenticateCredentials(body.Email, body.Password)
		if err != nil {
			log.Println(err)
			http.Error(w, "Unauthorized Access", http.StatusUnauthorized)
			return
		}
		// TODO remember to add role here
		if role == nil {
			t := ""
			role = &t
		}
		accessToken, err := service.SignAuthJWT(user.ID, *role)
		if err != nil {
			log.Println(err)
			http.Error(w, "Unauthorized Access", http.StatusInternalServerError)
			return
		}
		loginRes := loginResponse{AccessToken: accessToken}
		data, _ := json.Marshal(loginRes)
		_, _ = w.Write(data)
	})
}
