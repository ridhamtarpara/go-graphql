package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/vektah/gqlgen/handler"

	"ridham.me/jobs/graph"
	"ridham.me/jobs/graph/database"
)

func main() {
	app := &graph.App{}

	database.InitializeFirebase()

	queryHandler := cors.Default().Handler(http.HandlerFunc(handler.Playground("Job", "/query")))
	rootHandler := cors.Default().Handler(http.HandlerFunc(handler.GraphQL(graph.MakeExecutableSchema(app))))

	http.Handle("/", queryHandler)
	http.Handle("/query", rootHandler)

	fmt.Println("Listening on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
