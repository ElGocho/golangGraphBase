package services

import (
	"gorm.io/gorm"

	"sa_web_service/internal/models"
	"sa_web_service/internal/models/consts"
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

