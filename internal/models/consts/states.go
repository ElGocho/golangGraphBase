package consts

//Enumeración de estados
const (
	StateActUser State = "user_active"
	StateDisaUser State = "user_disable"

	StateActArticle State = "article_active"
	StateDisaArticle State = "article_disable"

	StateActReceipt	State = "receipt_active"
	StateDisaReceipt State = "receipt_disable"

	StateActReceiptArticle State = "receipt_article_active"
	StateAcceptedReceiptArticle	State = "receipt_article_accepted"



	StateActPrice	bool = true
	StateDisaPrice bool = false
)

 
func (s State) Name() string {
	switch(s){
	case	StateActUser:	return	"Activo"
	case	StateDisaUser:	return	"Deshabilitado"
	case	StateActArticle:	return	"Activo"
	case	StateDisaArticle:	return	"Deshabilitado"
	case	StateActReceipt:	return	"Activo"
	case	StateDisaReceipt:	return	"Deshabilitado"
	case	StateActReceiptArticle:	return	"Activo"
	case	StateAcceptedReceiptArticle:	return	"Aceptado"
	}
	
	return ""
}
