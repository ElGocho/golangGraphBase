package seeds

import (
	"sa_web_service/internal/models"
)

func GetAll() (resp []interface{}) {
	resp = append(resp, getWTables())
	resp = append(resp, getWCountries())
	resp = append(resp, getWRoles())

	return
}

func getWRoles() []*models.WRole{
	return []*models.WRole{
		&models.WRole{
			Name: "Administrador",
			Code: "admin",
			Status: true,
		},
		&models.WRole{
			Name: "Cliente",
			Code: "client",
			Status: true,
		},
		&models.WRole{
			Name: "Vendedor",
			Code: "seller",
			Status: true,
		},
	}
}

//Por pais se debe agregar los datos que dependen de determinados paises, ejp: Lenguajes y Monedas.
func getWCountries() []*models.WCountry{
	return []*models.WCountry{
		&models.WCountry{
			Name: "Venezuela",
			Code: "VEN",
			Status: true,
			Currencies: []*models.WCurrency{
				&models.WCurrency{
					Name:	"Bs",
					Code:	"VES",
					Status:	true,
				},
			},
			Languages:	[]*models.WLanguage{
				&models.WLanguage{
					Name: "Español venezuela",
					Code:	"es-VE",
					Status: true,
				},
			},
		},
	}
}

func getWTables() []*models.WTable{
	return []*models.WTable{
		&models.WTable{
			Name: "Usuarios",
			Code: "users",
			Status: true,
		},
		&models.WTable{
			Name: "Artículos",
			Code: "Articles",
			Status: true,
		},
		&models.WTable{
			Name: "Precios",
			Code: "prices",
			Status: true,
		},
		&models.WTable{
			Name: "Recibos",
			Code: "receipts",
			Status: true,
		},
		&models.WTable{
			Name: "Artículos de los Recibos",
			Code: "receipts_articles",
			Status: true,
		},
	}
}
