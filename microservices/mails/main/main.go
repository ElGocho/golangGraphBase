package main

import (
	"log"
	"fmt"
	"time"

	"gorm.io/gorm"
	"github.com/streadway/amqp"

	"sa_web_service/internal/helpers"
	"sa_web_service/internal/models/consts"

	mailHelpers "sa_web_service/microservices/mails/helpers"
	"sa_web_service/microservices/mails/controllers"
)


func main(){
	log.Println("Iniciando Microservicio: ", consts.MicroserviceMails)

	env := &helpers.ENV{}

	env.Load(".env")
	mailHelpers.Env.Load(".env")

	conn, ch, err := helpers.MQChannel(env.MQ_USER, env.MQ_PASSWORD, env.MQ_HOST, env.MQ_PORT)

	helpers.FailOnError(err, "Error al abrir el canal con rabbitMQ")

	defer conn.Close()
	defer ch.Close()

	queue,err := helpers.MQQueue(consts.QueueEmails, ch)

	helpers.FailOnError(err, fmt.Sprintf("Error al declarar cola %s en rabbitMQ", consts.QueueEmails))

	consumer,err := helpers.MQConsumer(ch, queue)

	helpers.FailOnError(err, fmt.Sprintf("Error al crear consumidor para la cola %s en rabbitMQ", consts.QueueEmails))

	db, err := helpers.InitDB(env.DB_USER, env.DB_PASSWORD, env.DB_HOST, env.DB_PORT, env.DB_NAME, env.DB_SSLMODE, env.LOG_MODE)

	helpers.FailOnError(err, "Error al iniciar la conexión con la BD")

	getMessageToQueue(consumer, db)
}


func getMessageToQueue(consumer <-chan amqp.Delivery, tx *gorm.DB){
	for delivery := range consumer{
		message := helpers.Message{}
		err := message.Deserialize(delivery.Body)

		if err != nil {
			log.Println("Error al deserializar el contenido del mensaje, se procedera a liberar el mesaje: ", delivery.MessageId)	
			delivery.Ack(true)
			continue
		}

		err, ack := controllers.Events(tx,message, delivery)

		if err != nil {
			log.Println(err)
			log.Println(ack)
		}

		delivery.Ack(ack)

		time.Sleep(1 * time.Second)
	}

}



