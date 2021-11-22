package consts

const (
	CountryVenezuela Country = "VEN"
)

func (c Country) Name() string{
	switch(c){
	case CountryVenezuela: return "Venezuela"
	}

	return ""
}
