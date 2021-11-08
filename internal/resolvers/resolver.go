package resolvers

import (
	"sa_web_service/graph/generated"
	"sa_web_service/internal/resolvers/queries"
	"sa_web_service/internal/resolvers/mutations"
	"sa_web_service/internal/models" 

	"gorm.io/gorm"
)

type Resolver struct{
	DB *gorm.DB
	ENV *models.ENV
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutations.MutationResolver{ r.DB, r.ENV} }

func (r *Resolver) Query() generated.QueryResolver { return &queries.QueryResolver{ r.DB, r.ENV} }

