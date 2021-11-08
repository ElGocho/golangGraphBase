package consts

const (
	CountryVenezuela country = "VEN"
)

func (c country) Name() string{
	switch(c){
	case CountryVenezuela: return "Venezuela"
	}

	return ""
}
