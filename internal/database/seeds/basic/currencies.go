package basic

import (
	"sa_web_service/internal/models"
	"sa_web_service/internal/models/const"
)

func GetWCurrenciesVenezuela() []*models.WCurrency{
	return []*models.WCurrency{
		&models.WCurrency{
			Name: cons.CurrencyVenezuela.Name(),
			Code: string(cons.CurrencyVenezuela),
			Status: true,
		},
	}
}
