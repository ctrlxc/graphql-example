package main

//go:generate go run gqlgen.go
//go:generate sqlboiler psql

import (
	"app/graph"
	"app/graph/generated"
	"app/loader"
	"app/repository"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"
const defaultDsn = "host=postgres user=graphql password=graphql dbname=graphql sslmode=disable"

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}

func run(_ []string) error {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dsn := os.Getenv("DSN")
	if dsn == "" {
		dsn = defaultDsn
	}

	repo, err := repository.New(dsn)
	if err != nil {
		return fmt.Errorf("failed to create repository: %+v", err)
	}

	resolver := graph.NewResolver(repo)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	middleware := func(next http.Handler) http.Handler {
		return loaderMiddleware(next, repo)
	}

	http.Handle("/", middleware(playground.Handler("GraphQL playground", "/query")))
	http.Handle("/query", middleware(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	return nil
}

// Middleware for attaching data loaders for GraphQL
func loaderMiddleware(next http.Handler, repo *repository.Repository) http.Handler {
	loaders := loader.NewLoaders(repo)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r.WithContext(loaders.Attach(r.Context())))
	})
}
