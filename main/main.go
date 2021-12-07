package main

import (
	"fmt"
	"sa_web_service/internal/database"
	"sa_web_service/internal/handlers"
	"sa_web_service/internal/handlers/middlewares"
	"sa_web_service/internal/helpers"
	"sa_web_service/internal/models/consts"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

func main() {
	env := &helpers.ENV{}

	env.Load(".env")

	db := getDB(env)

	autoFunc(db, env)

	channel, queues := loadQueues(env)

	fmt.Printf("[MAIN] Subiendo servidor en puerto %s\n",env.PORT)

	serverUp(db,env,channel,queues)
}

func serverUp(db *gorm.DB, env *helpers.ENV, channel *amqp.Channel, queues map[string]amqp.Queue){
	r := gin.Default()
	
	r.Use(middlewares.GinContextToContext())
	r.Use(middlewares.AuthMiddleware())

	r.POST("/query", handlers.GraphQL(db,env, channel, queues))
	r.GET("/", handlers.Playground())
	
	r.Run()
}

func getDB (env *helpers.ENV) (db *gorm.DB){
	db, err :=  helpers.InitDB(env.DB_USER, env.DB_PASSWORD, env.DB_HOST, env.DB_PORT, env.DB_NAME, env.DB_SSLMODE, env.LOG_MODE)

	helpers.FailOnError(err, "Error al abrir la conexión con la BD")

	return
}

func autoFunc(db *gorm.DB, env *helpers.ENV) {

	if env.AUTO_MIGRATE {
		database.ExecAll(db)
	}

	if env.GIN_MODE == "release"{
		gin.SetMode(gin.ReleaseMode)
	}

}

func loadQueues(env *helpers.ENV) (*amqp.Channel, map[string]amqp.Queue) {
	_, ch, err := helpers.MQChannel(env.MQ_USER, env.MQ_PASSWORD, env.MQ_HOST, env.MQ_PORT)

	helpers.FailOnError(err, "Error al abrir el canal con rabbitMQ")

	queues := make(map[string]amqp.Queue)

	queues[consts.QueueHello] = declareQueue(consts.QueueHello, ch)
	queues[consts.QueueEmails] = declareQueue(consts.QueueEmails, ch)

	return ch, queues
}

func declareQueue(qName string, ch *amqp.Channel) (amqp.Queue){
	q, err := helpers.MQQueue(qName, ch)

	helpers.FailOnError(err, fmt.Sprintf("Error al declarar cola %s en rabbitMQ", qName))

	return q
}
