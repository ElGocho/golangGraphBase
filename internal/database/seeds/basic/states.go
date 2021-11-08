package basic

import (
	"sa_web_service/internal/models"
	"sa_web_service/internal/models/consts"
)

func GetWStatesUser() []*models.WState {
	return []*models.WState{
		&models.WState{
			Name: consts.StateActUser.Name(),
			Code: string(consts.StateActUser),
			Status: true,
		},
	}
}

