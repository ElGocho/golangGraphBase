package handlers

import (
	"sa_web_service/graph/generated"
	"sa_web_service/internal/resolvers"
	"sa_web_service/internal/models"
	"sa_web_service/internal/handlers/middlewares"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func GraphQL(db *gorm.DB, env *models.ENV) gin.HandlerFunc{
	r := &resolvers.Resolver{
		db,
		env,
	}
	
	config := generated.Config{Resolvers:r}

	config.Directives.IsAuthenticated = middlewares.IsAuthenticated

	h := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func Playground() gin.HandlerFunc{
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
