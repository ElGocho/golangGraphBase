package services

import (
	"gorm.io/gorm"

	gql "sa_web_service/graph/model"
	"sa_web_service/internal/models"
	"sa_web_service/internal/models/consts"
	"sa_web_service/internal/models/builders"
)

func CreateArticles(tx *gorm.DB, articles models.Articles) (models.Articles, error){
	db := models.NewSession(tx)

	err := configArticles(db, articles, consts.StateActArticle, consts.StateActPrice)

	if err != nil {
		return nil, err
	}

	err = articles.Create(db).Error

	return articles, err 
}

func Articles(tx *gorm.DB, qBuilder *gql.QueryBuilder, pBuilder *models.Builder) (resp models.Articles){
	builder := &models.Builder{}
	priority := models.Priority1
	db := models.NewSession(tx)	

	if qBuilder != nil{
		b := builders.ArticleFromGQL(qBuilder)

		builder.Merge(b, &priority)
	}

	builder.Merge(pBuilder, &priority)

	resp.Find(db, builder)

	return
}

func UpdateArticles(tx *gorm.DB, articles models.Articles) (models.Articles, error){
	db := models.NewSession(tx)

	builder := &models.Builder{
		Omit: []string { "SellerID", "TableID", "CreatedAt","Prices","FTableID","TableItemID"},
	}

	err := db.Transaction(func(db *gorm.DB) error{
		err := articles.Save(db, builder)
		
		return err
	})

	return articles, err
}


func configArticles(db *gorm.DB, articles models.Articles, state consts.State, priceState bool) error {

	for _, article := range articles{
		err := configArticle(db, article, state)

		if err != nil {
			return err
		}

		err = configPrices(db, article.Prices, article.TableID, priceState)

		if err != nil {
			return err
		}
	}

	return nil
}

func configArticle(db *gorm.DB, article *models.Article, state consts.State) error{
	builder := models.SelectITable(consts.TableArticles, "id")

	err := article.Table.Get(db, builder).Error

	if err != nil {
		return err
	}

	article.TableID = article.Table.ID

	builder = models.SelectIState(state, article.TableID, "id")

	err = article.State.Get(db, builder).Error

	if err != nil {
		return err
	}

	article.StateID = article.State.ID

	return nil
}

