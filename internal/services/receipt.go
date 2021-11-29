package services

import (
	"gorm.io/gorm"

	gql "sa_web_service/graph/model"
	"sa_web_service/internal/models"
	"sa_web_service/internal/models/consts"
	"sa_web_service/internal/models/builders"
)

func CreateReceipts(tx *gorm.DB, receipts models.Receipts) (models.Receipts, error){
	db := models.NewSession(tx)

	err := configReceipts(db, receipts, consts.StateActReceipt, consts.StateActReceiptArticle)

	if err != nil {
		return nil, err
	}

	err = receipts.Create(db).Error

	return receipts, err 
}

func Receipts(tx *gorm.DB, qBuilder *gql.QueryBuilder, pBuilder *models.Builder) (resp models.Receipts){
	builder := &models.Builder{}
	priority := models.Priority1
	db := models.NewSession(tx)	

	if qBuilder != nil{
		b := builders.ReceiptFromGQL(qBuilder)

		builder.Merge(b, &priority)
	}

	builder.Merge(pBuilder, &priority)

	resp.Find(db, builder)

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

