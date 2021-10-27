package main

import (
	"fmt"
	"log"
	"net/http"
	"sa_web_service/graph/generated"
	"sa_web_service/internal/resolvers"
	"sa_web_service/internal/models"
	"sa_web_service/internal/database"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	env := &models.ENV{}

	env.Load(".env")

	port := env.PORT
	if port == "" {
		port = defaultPort
	}

	db := getDB(env)

	autoFunc(db, env)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: getResolver(db, env)}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getDB (env *models.ENV) (db *gorm.DB){
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", env.DB_HOST, env.DB_USER, env.DB_PASSWORD, env.DB_NAME, env.DB_PORT, env.DB_SSLMODE)

	db,err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return
}

func getResolver(db *gorm.DB, env *models.ENV) (r *resolvers.Resolver){
	r = &resolvers.Resolver{
		db,
		env,
	}

	return
}

func autoFunc(db *gorm.DB, env *models.ENV) {

	if env.AUTO_MIGRATE {
		database.ExecAll(db)
	}

}
