package mutations 

import (
	"context"

	gql "sa_web_service/graph/model"
	"sa_web_service/internal/models"
	sr "sa_web_service/internal/services"

)

func (r *MutationResolver) CreateArticles(ctx context.Context, input gql.CUArticleInput) ([]*models.Article, error) {
	articles, err := sr.CreateArticles(r.DB, input.Articles)	

	return articles, err
}

func (r *MutationResolver) UpdateArticles(ctx context.Context, input gql.CUArticleInput) ([]*models.Article, error) {
	articles, err := sr.UpdateArticles(r.DB, input.Articles)	

	return articles, err
}
