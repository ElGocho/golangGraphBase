package consts 

const (
	CurrencyVenezuela Currency = "VES"
)

func (c Currency) Name() string {
	switch(c){
	case CurrencyVenezuela: return "Bs"
	}

	return ""
}
