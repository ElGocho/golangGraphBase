package basic

import (
	"sa_web_service/internal/models"
	"sa_web_service/internal/models/consts"
)

func GetWLanguagesVenezuela() []*models.WLanguage{
	return []*models.WLanguage{
		&models.WLanguage{
			Name: consts.LanguageVenezuela.Name(),
			Code: string(consts.LanguageVenezuela),
			Status: true,
		},
	}
}
