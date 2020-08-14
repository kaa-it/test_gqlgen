package main

import (
	"net/http"
	graph "test_gqlgen/graph/generated"
	"test_gqlgen/resolvers"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
)

func main() {
	addr := ":4545"

	router := mux.NewRouter()
	rootResolver := resolvers.Resolver{}
	config := graph.Config{Resolvers: &rootResolver}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(config))

	router.Handle("/playground", playground.Handler("Test GQLGEN", "/query"))
	router.Handle("/query", srv)

	http.ListenAndServe(addr, router)
}
