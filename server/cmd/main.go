package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"server/internal/graph"
	"server/internal/graph/generated"

	"server/internal/system/database/psql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	// Клиент для psql
	repo, err := psql.New(context.Background(), "postgres://module12_task05:module12_task05@localhost:5432/module12_task05?sslmode=disable&connect_timeout=5")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Hello from task 05 sever!")
	////////////////

	////////////////////
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		Repo: repo,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
