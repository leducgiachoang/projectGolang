package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/leducgiachoang/app-gorm/generated"
	"github.com/leducgiachoang/app-gorm/graphql"
	"github.com/leducgiachoang/app-gorm/database"
	"github.com/leducgiachoang/app-gorm/app"
	"github.com/leducgiachoang/app-gorm/migrations"
	
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	if db, err := database.InitDB(); err != nil {
		panic(err)
	} else {
		app.DB = db
		log.Printf("Done\n\n")
	}

	migrations.Migrate(app.DB)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graphql.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
