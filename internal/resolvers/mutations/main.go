package mutations 

import (
	"context"
	"gorm.io/gorm"

	"sa_web_service/internal/models/consts" 
	"sa_web_service/internal/helpers" 
	sr "sa_web_service/internal/services" 
	gql "sa_web_service/graph/model"
	"github.com/streadway/amqp"
)

type MutationResolver struct{
	DB *gorm.DB
	ENV *helpers.ENV
	MQChannel *amqp.Channel
	Queues map[string]amqp.Queue
}

func (r *MutationResolver) Ping2(ctx context.Context) (bool, error) {
	return true, nil
}

func (r *MutationResolver) Hello(ctx context.Context) (*string, error) {
	body := "Hello World"

	err := helpers.MQPublish(r.MQChannel,r.Queues[consts.QueueHello], body)

	if err != nil {
		return nil, err
	}

	return &body, nil
}

func (r *MutationResolver) Login(ctx context.Context, input gql.LoginInput) (*gql.LoginRequest, error){
	db := helpers.NewSession(r.DB)	

	token, err := sr.Login(db, input.Email, input.Password)

	if err != nil {
		return nil, err
	}

	resp := &gql.LoginRequest{
		Token: token,
	}

	return resp, nil
}
