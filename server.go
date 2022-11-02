package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/stephendatascientist/go-graphql-mongodb-api/database"
	"github.com/stephendatascientist/go-graphql-mongodb-api/graph"
	"github.com/stephendatascientist/go-graphql-mongodb-api/graph/generated"
	"github.com/stephendatascientist/go-graphql-mongodb-api/repositories"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	client := database.Connect()
	defer client.Disconnect(context.Background())
	if err != nil {
		log.Fatal("could not load the database")
	}

	CustomerRepository := repositories.CustomerRepository{Client: client}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		CustomerRepository: CustomerRepository,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
