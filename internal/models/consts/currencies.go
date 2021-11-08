package consts 

const (
	CurrencyVenezuela currency = "VES"
)

func (c currency) Name() string {
	switch(c){
	case CurrencyVenezuela: return "Bs"
	}

	return ""
}
