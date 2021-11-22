package middlewares

import (
	"context"
	"strings"

	"sa_web_service/internal/models/consts"
	"sa_web_service/internal/helpers"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gin-gonic/gin"
)

type authHeader struct {
	Authorization	string `header:"Authorization"`
}

func GinContextToContext() gin.HandlerFunc {
	return func(c *gin.Context){
		ctx := context.WithValue(c.Request.Context(), consts.GinContext, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context){
		h := authHeader{}	

		err := c.ShouldBindHeader(&h)

		if err != nil {
			c.Next()
			return
		}

		token := strings.Split(h.Authorization, "Bearer ")

		if len(token) < 2 {
			c.Next()
			return
		}

		ctx := context.WithValue(c.Request.Context(), consts.AuthorizationKey, token[1])
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func IsAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error){
	err := helpers.ValidAuthToken(ctx) 

	if err != nil {
		return nil, err
	}

	return next(ctx)
}
