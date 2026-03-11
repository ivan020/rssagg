package main;

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"sql"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	_"github.com/lib/pq"
);

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env");

	portString := os.Getenv("PORT");
	if portString == "" {
		log.Fatal("PORT is not found in the env");
	}

	// import the database connection
	dbURL := os.Getenv("DB_URL");
	if dbURL == "" {
		log.Fatal("dbURL is not found in the env");
	}

	conn, err := sql.Open("postgres", dbURL);
	if err != err {
		log.Fatal("Can't connect to the database");
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter();

	// browser will allow request unless they meet these requirements 
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET","POST","PUT","DELETE","OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}));

	
	v1Router := chi.NewRouter();
	v1Router.Get("/healthz", handlerReadiness); // this is the handler
	v1Router.Get("/err", handlerErr);
	v1Router.Post("/users", apiCfg.handlerCreateUser); // so that the handler can create user in the database
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser)); 
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed)); 

	router.Mount("/v1", v1Router);


	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	};
	
	log.Printf("Server starting on port %v", portString);
	err := srv.ListenAndServe();
	if err != nil {
		log.Fatal(err);
	}

	fmt.Println("Port:", portString);
}
