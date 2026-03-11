package main;

import (
 "net/http" 
 "github.com/google/uuid"
)


func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL string `json:"url"`

	}
	decoder := json.NewDecoder(r.Body);
	decoder.Decode();
	
	params := parameters{};
	err := decoder.Decode(&params);
	if err != nil {
		respondWithError(w, 400,fmt.Sprintf("Error parsing json: %v", err));
		return;
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedA: time.Now().UTC(),
		Name: params.Name,
		Url: params.URL,
		UserID: feed.ID,
	});

	if err != nil {
		respondWithError(w, 400,fmt.Sprintf("Couldn't create user: %v", err));
		return;
	}

	respondWithJson(w, 201, feed);

}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, feed database.Feed) {
	respondWithJson(w, 200, databaseFeedToFeed(feed));
}



