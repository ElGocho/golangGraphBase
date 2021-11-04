package mutations 

import (
	"context"

	gql "sa_web_service/graph/model"
	sr	"sa_web_service/internal/services"
	"sa_web_service/internal/models"

)

func (r *MutationResolver) Register(ctx context.Context, input gql.RegisterInput) (*models.User, error) {
	_, err := sr.RegisterUser(r.DB,input.User)

	return input.User, err 
}
