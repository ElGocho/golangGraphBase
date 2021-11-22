package consts

//Enumeración de estados
const (
	StateActUser State = "user_active"
	StateDisaUser State = "user_disable"

	StateActArticle State = "article_active"
	StateDisaArticle State = "article_disable"

	StateActPrice	bool = true
	StateDisaPrice bool = false
)

 
func (s State) Name() string {
	switch(s){
	case	StateActUser:	return	"Activo"
	case	StateDisaUser:	return	"Deshabilitado"
	case	StateActArticle:	return	"Activo"
	case	StateDisaArticle:	return	"Deshabilitado"
	}
	
	return ""
}
