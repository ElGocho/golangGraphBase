package queries

import (
	"context"
	
	"sa_web_service/internal/helpers"
	gql "sa_web_service/graph/model"
	"sa_web_service/internal/models"
	sr "sa_web_service/internal/services"
)

func (r *QueryResolver) Articles(ctx context.Context, input *gql.QueryInput) ([]*models.Article, error){
	var qBuilder *gql.QueryBuilder 
	db := helpers.NewSession(r.DB)	

	if input != nil {
		qBuilder = input.Builder
	}

	resp := sr.Articles(db, qBuilder, nil)

	return resp, nil
}



