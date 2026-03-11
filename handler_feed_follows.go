package main;

import (
 "net/http" 
 "github.com/google/uuid"
)


func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID string `json:"feed_id"`

	}
	decoder := json.NewDecoder(r.Body);
	decoder.Decode();
	
	params := parameters{};
	err := decoder.Decode(&params);
	if err != nil {
		respondWithError(w, 400,fmt.Sprintf("Error parsing json: %v", err));
		return;
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedA: time.Now().UTC(),
		UserID: feed.ID,
		FeedID: params.FeedID,
	});

	if err != nil {
		respondWithError(w, 400,fmt.Sprintf("Couldn't create feed follow: %v", err));
		return;
	}

	respondWithJson(w, 201, databaseFeedFollowToFeedFollow( feedFollow ));
}


func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID);

	if err != nil {
		respondWithError(w, 400,fmt.Sprintf("Couldn't get feed follows: %v", err));
		return;
	}

	respondWithJson(w, 201, databaseFeedFollowsToFeedFollows( feedFollows ));
}


