package main

import (
	"fmt"
	"net/http"

	"github.com/KartikeChadha/feedWatch/internal/auth"
	"github.com/KartikeChadha/feedWatch/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apicfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth Error: %v", err))
			return
		}
		user, err := apicfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}
		handler(w, r, user)
	}
}
