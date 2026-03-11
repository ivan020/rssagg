package main;

import (
	"net/http"

	"github.com/ivan020/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User) 

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HanderFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
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

		handler(w, r, user);

	}
}
