package services

import (
	"gorm.io/gorm"

	"sa_web_service/internal/models"
	"sa_web_service/internal/models/consts"
)


func configReceiptArticles(db *gorm.DB, receiptArticles models.ReceiptArticles, state consts.State) error {

	for _, receiptArticle := range receiptArticles{
		err := configReceiptArticle(db, receiptArticle, state)

		if err != nil {
			return err
		}
	}

	return nil
}


func configReceiptArticle(db *gorm.DB, receiptArticle *models.ReceiptArticle, state consts.State) error{
	builder := models.SelectITable(consts.TableReceiptArticles, "id")

	err := receiptArticle.Table.Get(db, builder).Error

	if err != nil {
		return err
	}

	receiptArticle.TableID = receiptArticle.Table.ID

	builder = models.SelectIState(state, receiptArticle.TableID, "id")

	err = receiptArticle.State.Get(db, builder).Error

	if err != nil {
		return err
	}

	receiptArticle.StateID = receiptArticle.State.ID

	return nil
}
