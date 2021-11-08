package consts

//Enumeración de estados
const (
	StateActUser state = "active"
	StateDisaUser state = "disable"

)

 
func (s state) Name() string {
	switch(s){
	case StateActUser:	return	"Activo"
	case	StateDisaUser:	return	"Deshabilitado"
	}
	
	return ""
}
