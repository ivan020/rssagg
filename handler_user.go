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

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	
	apiKey, err := auth.GetAPIKey(r.Header);
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Auth error: %s", err));
		return;
	}
	
	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey); 
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get user: %s", err));
		return;
	}
	respondWithJson(w, 200, databaseUserToUser(user));

}
