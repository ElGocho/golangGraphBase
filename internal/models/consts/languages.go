package consts

const (
	LanguageVenezuela Language = "es-VE"
	LanguageEspaña Language = "es"
)

func (l Language) Name() string{
	switch(l){
	case LanguageVenezuela: return "Español venezuela"
	case LanguageEspaña: return "Español"
	}

	return ""
}

