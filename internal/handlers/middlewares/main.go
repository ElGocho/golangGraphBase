package middlewares

import (
	"context"

	"sa_web_service/internal/models/const"

	"github.com/gin-gonic/gin"
)

func GinContextToContext() gin.HandlerFunc {
	return func(c *gin.Context){
		ctx := context.WithValue(c.Request.Context(), cons.GinContext, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}



