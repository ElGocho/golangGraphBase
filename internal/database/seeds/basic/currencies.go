package basic

import (
	"sa_web_service/internal/models"
	"sa_web_service/internal/models/consts"
)

func GetWCurrenciesVenezuela() []*models.WCurrency{
	return []*models.WCurrency{
		&models.WCurrency{
			Name: consts.CurrencyVenezuela.Name(),
			Code: string(consts.CurrencyVenezuela),
			Status: true,
		},
	}
}
