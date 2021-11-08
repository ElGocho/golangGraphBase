package consts

const (
	LanguageVenezuela language = "es-VE"
	LanguageEspaña language = "es"
)

func (l language) Name() string{
	switch(l){
	case LanguageVenezuela: return "Español venezuela"
	case LanguageEspaña: return "Español"
	}

	return ""
}

