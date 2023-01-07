package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/midwhite/golang-web-server-sample/graphql-gateway/graph"
	"github.com/midwhite/golang-web-server-sample/graphql-gateway/userservice"
)

const defaultPort = "8080"

func main() {
	closer, err := userservice.Setup()
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
