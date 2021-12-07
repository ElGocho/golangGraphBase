package services

import (
	"errors"
	"gorm.io/gorm"

	"sa_web_service/internal/models" 
	"sa_web_service/internal/helpers" 
)


func Login(db *gorm.DB,email, password string) (string, error) {
	builder := &models.Builder{
		Select: "password,name",
		Where: []models.Where{
			models.Where{
				Condition: models.User{ Email: email},
			},
		},
	}
	
	user := &models.User{}

	err := user.Get(db, builder).Error

	if err != nil {
		return "", err
	}

	err = helpers.CompareHash(user.Password, password)

	if err != nil {
		return "", err
	}

	data := map[string]interface{} {
		"email": email,
		"name": user.Name,
	}

	token, err := helpers.GenerateToken(data)

	if err != nil {
		return "", errors.New("Error al crear el token")
	}

	return token, err 
}
