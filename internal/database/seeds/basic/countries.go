package basic

import (
	"sa_web_service/internal/models"
	"sa_web_service/internal/models/consts"
)


//Por pais se debe agregar los datos que dependen de determinados paises, ejp: Lenguajes y Monedas.
func GetWCountries() []*models.WCountry{
	return []*models.WCountry{
		&models.WCountry{
			Name: consts.CountryVenezuela.Name(),
			Code: string(consts.CountryVenezuela),
			Status: true,
			Currencies: GetWCurrenciesVenezuela(),
			Languages: GetWLanguagesVenezuela(),
		},
	}
}

