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
		respondWithError(w, 400,fmt.Sprintf("Error parsing json"));
		return;
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedA: time.Now().UTC(),
		Name: params.Name,
	});

	if err != nil {
		respondWithError(w, 400,fmt.Sprintf("Couldn't create user:", err));
		return;
	}
	respondWithJson(w, 200, struct{}{});

}
