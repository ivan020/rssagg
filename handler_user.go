package main;

import (
 "net/http" 
 "github.com/google/uuid"
)


func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body);
	decoder.Decode();
	
	params := parameters{};
	err := decoder.Decode(&params);
	if err != nil {
		respondWithError(w, 400,fmt.Sprintf("Error parsing json: %v", err));
		return;
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedA: time.Now().UTC(),
		Name: params.Name,
	});

	if err != nil {
		respondWithError(w, 400,fmt.Sprintf("Couldn't create user: %v", err));
		return;
	}

	respondWithJson(w, 201, databaseUserToUser(user));

}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, database.User) {
	respondWithJson(w, 200, databaseUserToUser(user));
}

func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		limit: 10,
	});

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get posts: %v", err));
		return;
	}

	respondWithJson(w, 200, datbasePostsToPosts(posts));

}

