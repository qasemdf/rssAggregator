package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello World")

	//Creating a running server on port 8080
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the env")
	}
	fmt.Println("Port:", portString)

	router := chi.NewRouter()

	//Making it so people can make requests from their browser
	//cors tells our server to allow the user to do whatever you want
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	//created a new router
	v1Router := chi.NewRouter()
	//connecting handlerReadiness function to the /ready path
	//only going to fire on get Request
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handleErr)

	router.Mount("/v1", v1Router)

	srv := http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	log.Printf("Server Starting on %v", portString)
	//Handles http requests and checks for errors and if so it kills the program
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
