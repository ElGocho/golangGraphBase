package consts

type Template string

const (
	Templates Template = "microservices/mails/templates/"
	TemplateMailWelcome Template = Templates+"mails/welcome.html"

	TemplateMailReceiptCreate Template = Templates+"mails/receiptCreate.html"
)


func (t Template) Name() string {

	switch(t){
	case Templates: return string(Templates)
	case TemplateMailWelcome: return "welcome.html"
	case TemplateMailReceiptCreate: return "receiptCreate.html"
	default:
		return ""
	}
}
