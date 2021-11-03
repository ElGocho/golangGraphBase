package helpers

import (
	"context"
	"fmt"

	"sa_web_service/internal/models/const"
	"github.com/gin-gonic/gin"
)

func GetFromContext(ctx context.Context, key string) (interface{}, error){
	value := ctx.Value(key)

	if value == nil {
		err := fmt.Errorf("could not retrieve context")
		return nil, err
	}

	return value, nil
}

func GetGinContext(ctx context.Context) (*gin.Context, error) {
	value, err := GetFromContext(ctx, cons.GinContext)

	if err != nil{
		return nil, err
	}

	gc, ok := value.(*gin.Context)

	if !ok {
		err = fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	
	return gc, nil
}



