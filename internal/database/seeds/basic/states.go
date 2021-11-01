package basic

import (
	"sa_web_service/internal/models"
	"sa_web_service/internal/models/const"
)

func GetWStatesUser() []*models.WState {
	return []*models.WState{
		&models.WState{
			Name: cons.StateActUser.Name(),
			Code: string(cons.StateActUser),
			Status: true,
		},
	}
}

