package models 

import (
	"fmt"
	"gorm.io/gorm"
	
	"sa_web_service/internal/models/consts"
)

func(r *repository) ReceiptCode(db *gorm.DB, model *Receipt) (err error){
	var count int64

	db.Model(model).Count(&count)

	model.Code = fmt.Sprintf("R-%04d",count)

	return
}

//carga el monto total de los Receipt según los prices que estan cargados en el objeto
func (r *repository)ReceiptAmountTotal(db *gorm.DB, model *Receipt) (err error){
	if model == nil || len(model.ReceiptArticles) == 0 {
		return
	}  

	var pricesId []string

	result := struct {
		Total	float64
	}{}

	for _,receiptArticle := range model.ReceiptArticles {
		pricesId = append(pricesId, receiptArticle.PriceID.String())
	}

	err = db.Table(string(consts.TablePrices)).
			Select("sum(amount) as Total").
			Where("id IN ?", pricesId).
			Scan(&result).
			Error
	
	if err != nil {
		return
	}

	model.AmountTotal = result.Total

	return
}
