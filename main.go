package main

import (
	"database/sql"
	"fmt"
	"github.com/nell209/AutumnRefactor/database"
	"github.com/nell209/AutumnRefactor/middleware"
	"github.com/nell209/AutumnRefactor/service"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/nell209/AutumnRefactor/graph"
	"github.com/nell209/AutumnRefactor/graph/generated"
)

const defaultPort = "8080"

//func startGoogleProfiler() {
//	cfg := profiler.Config{
//		Service:        "AutumnBackendRefactor",
//		ServiceVersion: "0.0.0",
//	}
//	if err := profiler.Start(cfg); err != nil {
//		panic(fmt.Errorf("could not start google profiler"))
//	}
//}

func main() {
	//startGoogleProfiler()
	// Don't forget the stripe stuff
	// Should add a decent logger or something here
	// Could create our own error logs part of the DB

	// TODO resolve this from .env
	db, err := database.CreateConnection("dev")
	if err != nil {
		panic("Could not create a database connection")
	}
	closeable, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("Failed to get closeable function for database but got error %v", err))
	}
	defer func(closeable *sql.DB) {
		err := closeable.Close()
		if err != nil {
			panic(fmt.Sprintf("Failed to close the database connection but got error %v", err))
		}
	}(closeable)

	database.Migrate(db)

	resolver := graph.BindServicesToResolver(service.InitializeServices(db))
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", middleware.AuthMiddleware(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
