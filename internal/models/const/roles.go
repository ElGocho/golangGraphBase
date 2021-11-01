package cons

const (
	RoleAdministrator	role = "administrator"
	RoleClient role = "client"
	RoleSeller role = "seller"
)

func (r role) Name() string{
	switch(r){
	case RoleAdministrator: return "Administrador"
	case RoleClient:	return "Cliente"
	case RoleSeller:	return "Vendedor"
	}

	return ""
}
