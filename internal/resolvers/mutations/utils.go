package mutations 

import (
	"context"
	"gorm.io/gorm"

	"sa_web_service/internal/models" 
	sr "sa_web_service/internal/services" 
	gql "sa_web_service/graph/model"
)

type MutationResolver struct{
	DB *gorm.DB
	ENV *models.ENV
}

func (r *MutationResolver) Ping2(ctx context.Context) (bool, error) {
	return true, nil
}

func (r *MutationResolver) Login(ctx context.Context, input gql.LoginInput) (*gql.LoginRequest, error){

	token, err := sr.Login(r.DB, input.Email, input.Password)

	if err != nil {
		return nil, err
	}

	resp := &gql.LoginRequest{
		Token: token,
	}

	return resp, nil
}
