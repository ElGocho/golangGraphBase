package basic

import (
	"sa_web_service/internal/models"
	"sa_web_service/internal/models/const"
)

func GetWLanguagesVenezuela() []*models.WLanguage{
	return []*models.WLanguage{
		&models.WLanguage{
			Name: cons.LanguageVenezuela.Name(),
			Code: string(cons.LanguageVenezuela),
			Status: true,
		},
	}
}
