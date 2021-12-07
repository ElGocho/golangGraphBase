package services

import (
	"fmt"
	"gorm.io/gorm"
	"github.com/streadway/amqp"

	"sa_web_service/internal/helpers"
	gql "sa_web_service/graph/model"
	"sa_web_service/internal/models"
	"sa_web_service/internal/models/consts"
	"sa_web_service/internal/models/builders"
)

func CreateReceipts(db *gorm.DB, receipts models.Receipts) (models.Receipts, error){
	err := configReceipts(db, receipts, consts.StateActReceipt, consts.StateActReceiptArticle)

	if err != nil {
		return nil, err
	}

	err = receipts.Create(db).Error

	if err != nil {
		return nil, err
	}

	return receipts, err 
}

func Receipts(db *gorm.DB, qBuilder *gql.QueryBuilder, pBuilder *models.Builder) (resp models.Receipts){
	builder := &models.Builder{}
	priority := models.Priority1

	if qBuilder != nil{
		b := builders.ReceiptFromGQL(qBuilder)

		builder.Merge(b, &priority)
	}

	builder.Merge(pBuilder, &priority)

	resp.Find(db, builder)

	return
}

func ReceiptsSendEmailToSellers(db *gorm.DB,receipts models.Receipts, eventName string, ch *amqp.Channel, q amqp.Queue) (err error){
	var mData []string

	sellers, err := receiptsLoadSellers(db, receipts)

	if err != nil {
		return err
	}

	for _, seller := range sellers{
		mData = append(mData, seller.ID.String())
	}

	m := helpers.Message{
		Data:	mData,
		Event: helpers.QueueEvent{
			Name: eventName,
		},
	}

	err = helpers.MQPublishMessage(ch,q, m)

	return
}

func ReceiptsSendEmailToClients(receipts models.Receipts, eventName string, ch *amqp.Channel, q amqp.Queue) (err error){
	var mData []string

	for _, receipt := range receipts {
		mData = append(mData, receipt.ClientID.String())
	}

	m := helpers.Message{
		Data:	mData,
		Event: helpers.QueueEvent{
			Name: eventName,
		},
	}

	err = helpers.MQPublishMessage(ch, q, m)

	return
}

func receiptsLoadSellers(db *gorm.DB, receipts models.Receipts) (users models.Users, err error){
	builder := &models.Builder{}	
	ids := []interface{}{}

	builders.GetUsersJoins(builder, []string{"receipt"})

	for _, receipt := range receipts {
		ids = append(ids, receipt.ID.String())
	}

	where := models.Where{
		Condition: fmt.Sprintf("%s.id in (?)", consts.TableReceipts), 
		Params: ids,
	}

	builder.Where = append(builder.Where, where)

	err = users.Find(db, builder).Error

	return 
}

func configReceipts(db *gorm.DB, receipts models.Receipts, state consts.State, receiptArticleState consts.State) error {

	for _, receipt := range receipts{
		err := configReceipt(db, receipt, state)

		if err != nil {
			return err
		}

		err = configReceiptArticles(db, receipt.ReceiptArticles, receiptArticleState)

		if err != nil {
			return err
		}


	}

	return nil
}


func configReceipt(db *gorm.DB, receipt *models.Receipt, state consts.State) error{
	builder := models.SelectITable(consts.TableReceipts, "id")

	err := receipt.Table.Get(db, builder).Error

	if err != nil {
		return err
	}

	receipt.TableID = receipt.Table.ID

	builder = models.SelectIState(state, receipt.TableID, "id")

	err = receipt.State.Get(db, builder).Error

	if err != nil {
		return err
	}

	receipt.StateID = receipt.State.ID

	return nil
}

