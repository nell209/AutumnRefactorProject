package restHandler

import (
	"log"
	"net/http"
)

// create an admin user
// create a password
// create a project
// create a couple tasks

func (h *restHandler) Init() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if err := h.service.Init(r.Context()); err != nil {
			log.Println(err)
			http.Error(w, "couldn't set password", http.StatusInternalServerError)
		}
		// Get User Branch
		w.Write([]byte("success"))
	})
}
