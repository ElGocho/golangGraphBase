package mutations 

import (
	"context"
	"gorm.io/gorm"

	"sa_web_service/internal/models" 
)

type MutationResolver struct{
	DB *gorm.DB
	ENV *models.ENV
}

func (r *MutationResolver) Ping2(ctx context.Context) (bool, error) {
	return true, nil
}
