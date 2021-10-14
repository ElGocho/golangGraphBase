package queries

import (
	"context"
)

type QueryResolver struct {}

func (r *QueryResolver) Ping(ctx context.Context) (bool, error) {
	return true, nil
}

