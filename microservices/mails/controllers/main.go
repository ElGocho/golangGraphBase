package controllers

import (
	"fmt"
	"gorm.io/gorm"
	"github.com/streadway/amqp"

	"sa_web_service/internal/helpers"
	"sa_web_service/internal/models/consts"
	"sa_web_service/microservices/mails/events"
)


func Events(tx *gorm.DB, message helpers.Message, delivery amqp.Delivery) (err error, ack bool){
		db := helpers.NewSession(tx)

		switch(message.Event.Name) {
		case consts.QueueEmailsEventWelcome:
			err = events.Welcome(db, message)

			if err != nil {
				ack = false
			}

			break
		case consts.QueueEmailsEventReceiptCreate:
			err = events.ReceiptsCreate(db, message)

			if err != nil {
				ack = false
			}

			break
		default:
			err = fmt.Errorf("Evento no encontrado: %s, se procedera a liberar el mensaje: %s", message.Event.Name, delivery.MessageId)
			ack = true 
		}

		return
}
