package main

import (
	"log"
	"net/http"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/vektah/gqlgen/handler"

	"ridham.me/jobs/graph"
	"ridham.me/jobs/graph/database"
)

func main() {
	// Load environment variable from .env variable
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize firebase database
	database.InitializeFirebase()

	// Create graph app and start the server
	app := &graph.App{}
	queryHandler := cors.Default().Handler(http.HandlerFunc(handler.Playground("Job", "/query")))
	rootHandler := cors.Default().Handler(http.HandlerFunc(handler.GraphQL(graph.MakeExecutableSchema(app))))

	http.Handle("/", queryHandler)
	http.Handle("/query", rootHandler)

	fmt.Println("Listening on http://localhost",os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), nil))
}
