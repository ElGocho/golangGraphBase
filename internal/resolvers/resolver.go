package resolvers

import (
	"sa_web_service/graph/generated"
	"sa_web_service/internal/resolvers/queries"
	"sa_web_service/internal/resolvers/mutations"
)

type Resolver struct{}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutations.MutationResolver{} }

func (r *Resolver) Query() generated.QueryResolver { return &queries.QueryResolver{} }

