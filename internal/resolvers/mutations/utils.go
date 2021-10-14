package mutations 

import (
	"context"
)

type MutationResolver struct{ }

func (r *MutationResolver) Ping2(ctx context.Context) (bool, error) {
	return true, nil
}
