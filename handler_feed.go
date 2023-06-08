package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/KartikeChadha/feedWatch/internal/database"
	"github.com/google/uuid"
)

func (apiCfg apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error parsing JSON:%v", err))
		return
	}
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		log.Fatal("Can't create feed", err)
		return
	}
	respondWithJSON(w, 201, databaseFeedToFeed(feed))

}

func (apiCfg apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {

	feed, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		log.Fatal("Can't get feed", err)
		return
	}
	respondWithJSON(w, 201, databaseFeedsToFeeds(feed))

}
