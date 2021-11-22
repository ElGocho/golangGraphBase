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
		&models.WState{
			Name: consts.StateDisaUser.Name(),
			Code: string(consts.StateDisaUser),
			Status: true,
		},
	}
}


func GetWStatesArticle() []*models.WState{
	return []*models.WState{
		&models.WState{
			Name: consts.StateActArticle.Name(),
			Code: string(consts.StateActArticle),
			Status: true,
		},
		&models.WState{
			Name: consts.StateDisaArticle.Name(),
			Code: string(consts.StateDisaArticle),
			Status: true,
		},
	}

}
