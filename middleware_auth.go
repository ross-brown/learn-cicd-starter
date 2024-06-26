package main

import (
	"net/http"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
	"github.com/bootdotdev/learn-cicd-starter/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

// looks for API key and retrives user from DB with key
func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Couldn't find api key")
			return
		}

		user, err := cfg.DB.GetUser(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusNotFound, "Couldn't get user")
			return
		}

		handler(w, r, user)
	}
}
