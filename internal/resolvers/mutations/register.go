package mutations 

import (
	"context"

	gql "sa_web_service/graph/model"
	tf	"sa_web_service/internal/resolvers/transformations"
	sr	"sa_web_service/internal/services"

)

func (r *MutationResolver) Register(ctx context.Context, input gql.RegisterInput) (*gql.User, error) {
	user,err := tf.UserInputToUserModel(input.User)	

	if err != nil {
		return nil, err
	}

	user, err = sr.RegisterUser(r.DB,user)

	if err != nil {
		return nil, err
	}

	resp, err := tf.UserModelToUserGQL(user)

	return resp, nil
}
