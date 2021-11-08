package queries

import (
	"context"
	
	gql "sa_web_service/graph/model"
	"sa_web_service/internal/models"
	sr "sa_web_service/internal/services"
)

func (r *QueryResolver) Users(ctx context.Context, input *gql.QueryInput) ([]*models.User, error){
	var qBuilder *gql.QueryBuilder 

	if input != nil {
		qBuilder = input.Builder
	}

	resp := sr.Users(r.DB, qBuilder, nil)

	return resp, nil
}


