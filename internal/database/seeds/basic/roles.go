package basic

import (
	"sa_web_service/internal/models"
	"sa_web_service/internal/models/consts"
)

func GetWRoles() []*models.WRole{
	return []*models.WRole{
		&models.WRole{
			Name: consts.RoleAdministrator.Name(),
			Code: string(consts.RoleAdministrator),
			Status: true,
		},
		&models.WRole{
			Name: consts.RoleClient.Name(),
			Code: string(consts.RoleClient),
			Status: true,
		},
		&models.WRole{
			Name: consts.RoleSeller.Name(),
			Code: string(consts.RoleSeller),
			Status: true,
		},
	}
}
