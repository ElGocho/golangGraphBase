package migrations

import (
	"sa_web_service/internal/models"
)


func GetAll() (resp []interface{}){
	resp = append(resp, &models.ITable{})
	resp = append(resp, &models.ICountry{})
	resp = append(resp, &models.ICurrency{})
	resp = append(resp, &models.ILanguage{})
	resp = append(resp, &models.IRole{})
	resp = append(resp, &models.IState{})

	resp = append(resp, &models.Article{})
	resp = append(resp, &models.User{})
	resp = append(resp, &models.Price{})
	resp = append(resp, &models.Receipt{})
	resp = append(resp, &models.ReceiptArticle{})
	resp = append(resp, &models.StatusDate{})

	return
}
