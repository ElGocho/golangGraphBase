package services

import (
	"gorm.io/gorm"

	"sa_web_service/internal/models"
)



func configPrices(db *gorm.DB, prices []*models.Price,fTable uint, status bool) error {
	
	for _, elem := range prices {
		err := configPrice(db,elem,fTable,status)

		if err != nil {
			return err
		}
	}

	return nil
}

func configPrice(db *gorm.DB, price *models.Price,fTable uint, status bool) error {
	price.FTableID = fTable

	price.Status = status

	return nil
}

