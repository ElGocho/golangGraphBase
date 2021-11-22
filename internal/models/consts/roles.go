package consts

const (
	RoleAdministrator	Role = "administrator"
	RoleClient Role = "client"
	RoleSeller Role = "seller"
)

func (r Role) Name() string{
	switch(r){
	case RoleAdministrator: return "Administrador"
	case RoleClient:	return "Cliente"
	case RoleSeller:	return "Vendedor"
	}

	return ""
}
