package mutations 

import (
	"context"

	"sa_web_service/internal/helpers"
	gql "sa_web_service/graph/model"
	"sa_web_service/internal/models"
	sr "sa_web_service/internal/services"

)

func (r *MutationResolver) CreateArticles(ctx context.Context, input gql.CUArticleInput) ([]*models.Article, error) {
	db := helpers.NewSession(r.DB)

	articles, err := sr.CreateArticles(db, input.Articles)	

	return articles, err
}

func (r *MutationResolver) UpdateArticles(ctx context.Context, input gql.CUArticleInput) ([]*models.Article, error) {
	db := helpers.NewSession(r.DB)

	db = db.Begin()

	articles, err := sr.UpdateArticles(db, input.Articles)	

	if err != nil {
		db.Rollback()
		return nil, err
	}

	db.Commit()
	return articles, nil
}
