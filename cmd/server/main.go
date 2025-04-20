package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jukemori/timeline-generator/graph/generated"
	"github.com/jukemori/timeline-generator/graph/resolver"
	"github.com/jukemori/timeline-generator/internal/database"
	"github.com/jukemori/timeline-generator/internal/openai"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	// Initialize the database connection
	database.InitDB()
	
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	openaiClient := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &resolver.Resolver{
			OpenAIClient: openaiClient,
		},
	}))

	// Setup CORS
	allowOrigins := strings.Split(os.Getenv("ALLOW_ORIGINS"), ",")
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   allowOrigins,
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		MaxAge:           60 * 60, // 1 hour in seconds
	})

	// Create a new router
	mux := http.NewServeMux()
	
	// Add the handlers with CORS middleware
	mux.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	mux.Handle("/graphql", corsHandler.Handler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}