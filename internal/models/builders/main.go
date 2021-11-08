package builders

import (
	gql "sa_web_service/graph/model"
	"sa_web_service/internal/models"
)

func getPagination(builder *models.Builder, pagination *gql.Pagination) *models.Builder {
	if pagination == nil {
		return builder
	}

	builder.Pagination = &models.Pagination{}

	if pagination.Page != nil {
		builder.Pagination.Page = pagination.Page
	}

	if pagination.Skip != nil {
		builder.Pagination.Skip = pagination.Skip
	}

	if pagination.Take != nil {
		builder.Pagination.Take = pagination.Take
	}

	return builder
}

func getOrder(builder *models.Builder, order *string) *models.Builder{
	if order == nil {
		return builder
	}

	builder.Order = order

	return builder
}

func getGroup(builder *models.Builder, group *string) *models.Builder{
	if group == nil {
		return builder
	}

	builder.Group = group

	return builder
}

