package queries

import (
	"context"

	"gorm.io/gorm"

	"sa_web_service/internal/models"
)

type QueryResolver struct {
	DB *gorm.DB
	ENV *models.ENV
}

func (r *QueryResolver) Ping(ctx context.Context) (bool, error) {
	return true, nil
}
