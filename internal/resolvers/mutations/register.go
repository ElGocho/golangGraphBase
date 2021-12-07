package mutations 

import (
	"context"

	"sa_web_service/internal/helpers"
	gql "sa_web_service/graph/model"
	sr	"sa_web_service/internal/services"
	"sa_web_service/internal/models"
	"sa_web_service/internal/models/consts"

)

func (r *MutationResolver) Register(ctx context.Context, input gql.RegisterInput) (*models.User, error) {
	db := helpers.NewSession(r.DB)	

	user, err := sr.RegisterUser(db,input.User)

	if err != nil {
		return nil, err
	}

	err = sr.RegisterSendEmail(db, user, consts.QueueEmailsEventWelcome, r.MQChannel, r.Queues[consts.QueueEmails])

	if err != nil {
		return nil, err
	}

	return user, nil
}
