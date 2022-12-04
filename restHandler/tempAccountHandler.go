package restHandler

import (
	"encoding/json"
	"net/http"
)

type tempAccountCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type tempAccountResponse struct {
	AccessToken string `json:"accessToken"`
}

func (h *restHandler) AuthenticateTempAccount() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		decoder := json.NewDecoder(r.Body)
		var body tempAccountCredentials
		err := decoder.Decode(&body)
		if err != nil {
			http.Error(w, "error parsing request body", http.StatusBadRequest)
		}

	})
}
