package resolvers

import (
	"context"
)

func (r *queryResolver) Ping(ctx context.Context) (bool, error) {
	return true, nil
}


