package services

import (
	"gorm.io/gorm"
	"github.com/streadway/amqp"

	"sa_web_service/internal/helpers"
	gql "sa_web_service/graph/model"
	"sa_web_service/internal/models"
	"sa_web_service/internal/models/consts"
	"sa_web_service/internal/models/builders"
)

func RegisterUser(db *gorm.DB, user *models.User) (*models.User, error){
	err := configUser(db, user, consts.StateActUser)
	
	if err != nil {
		return nil, err
	}

	err = user.Create(db).Error

	return user, err
}

func Users(db *gorm.DB, qBuilder *gql.QueryBuilder, pBuilder *models.Builder) (resp models.Users){
	builder := &models.Builder{}
	priority := models.Priority1

	if qBuilder != nil{
		b := builders.UserFromGQL(qBuilder)

		builder.Merge(b, &priority)
	}

	builder.Merge(pBuilder, &priority)

	resp.Find(db, builder)

	return
}

func RegisterSendEmail(db *gorm.DB, user *models.User, eventName string, ch *amqp.Channel, q amqp.Queue) (err error){
	m := helpers.Message{
		Data:	user.ID.String(),
		Event: helpers.QueueEvent{
			Name: eventName,
		},
	}

	err = helpers.MQPublishMessage(ch, q, m)

	return
}

func configUser(db *gorm.DB, user *models.User, state consts.State) error {
	builder := models.SelectITable(consts.TableUsers, "id")

	err := user.Table.Get(db, builder).Error

	if err != nil {
		return err
	}

	user.TableID = user.Table.ID

	builder = models.SelectIState(state, user.TableID, "id")

	err = user.State.Get(db, builder).Error

	if err != nil {
		return err
	}

	user.StateID = user.State.ID

	return nil
}
