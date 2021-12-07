package queries

import (
	"context"

	"gorm.io/gorm"
	"github.com/streadway/amqp"

	"sa_web_service/internal/models/consts" 
	"sa_web_service/internal/helpers" 
)

type QueryResolver struct {
	DB *gorm.DB
	ENV *helpers.ENV
	MQChannel *amqp.Channel
	Queues map[string]amqp.Queue
}

func (r *QueryResolver) Ping(ctx context.Context) (bool, error) {
	return true, nil
}

func (r *QueryResolver) Hello(ctx context.Context) (*string, error) {
	body := "Hello World"

	err := helpers.MQPublish(r.MQChannel,r.Queues[consts.QueueHello], body)

	if err != nil {
		return nil, err
	}

	return &body, nil
}

