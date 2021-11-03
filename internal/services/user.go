package services

import (
	"gorm.io/gorm"

	"sa_web_service/internal/models"
	"sa_web_service/internal/models/const"
)

func RegisterUser(tx *gorm.DB, user *models.User) (*models.User, error){
	db := models.NewSession(tx)

	builder := &models.Builder{
		Select: "id",
		Where: []models.Where{
			{
				Condition: models.ITable{ Code: string(cons.TableUsers)},
			},
		},
	}

	err := user.Table.Get(db, builder).Error

	if err != nil {
		return nil, err
	}

	user.TableID = user.Table.ID

	builder = &models.Builder{
		Select: "id",
		Where: []models.Where{
			{
				Condition: models.IState{ 
					Code: string(cons.StateActUser),
					TableID: user.TableID,
				},
			},
		},
	}

	err = user.State.Get(db, builder).Error

	if err != nil {
		return nil, err
	}

	user.StateID = user.State.ID

	err = user.Create(db).Error

	return user, err
}
