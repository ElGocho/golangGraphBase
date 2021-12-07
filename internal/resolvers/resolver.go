package resolvers

import (
	"sa_web_service/graph/generated"
	"sa_web_service/internal/resolvers/queries"
	"sa_web_service/internal/resolvers/mutations"
	"sa_web_service/internal/helpers" 

	"gorm.io/gorm"
	"github.com/streadway/amqp"
)

type Resolver struct{
	DB *gorm.DB
	ENV *helpers.ENV
	MQChannel *amqp.Channel
	Queues map[string]amqp.Queue
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutations.MutationResolver{ r.DB, r.ENV, r.MQChannel, r.Queues} }

func (r *Resolver) Query() generated.QueryResolver { return &queries.QueryResolver{ r.DB, r.ENV, r.MQChannel, r.Queues} }

