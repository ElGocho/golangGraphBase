package basic

import (
	"sa_web_service/internal/models"
	"sa_web_service/internal/models/const"
)

func GetWRoles() []*models.WRole{
	return []*models.WRole{
		&models.WRole{
			Name: cons.RoleAdministrator.Name(),
			Code: string(cons.RoleAdministrator),
			Status: true,
		},
		&models.WRole{
			Name: cons.RoleClient.Name(),
			Code: string(cons.RoleClient),
			Status: true,
		},
		&models.WRole{
			Name: cons.RoleSeller.Name(),
			Code: string(cons.RoleSeller),
			Status: true,
		},
	}
}
