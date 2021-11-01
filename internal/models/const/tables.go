package cons

//Enumeración de Tablas
const (
	TableTables table = "i_tables"
	TableCountries table = "i_countries"
	TableLanguages table = "i_languages"
	TableStates table = "i_states"
	TableRoles	table	=	"i_roles"
	TableCurrencies table = "i_currencies"
	TableUsers table = "users"
	TableArticles table = "articles"
	TablePrices table = "prices"
	TableReceipts table = "receipts"
	TableReceiptArticles table = "receipt_articles"

	TableStatusDate table = "status_date"
)

func (t table) Name() string{
	switch(t){
	case TableTables: return "Tablas"
	case TableCountries: return "Paises"
	case TableLanguages: return "Lenguajes"
	case TableStates: return "Estatus"
	case TableRoles: return "Roles"
	case TableCurrencies: return "Monedas"
	case TableUsers: return "Usuarios"
	case TableArticles: return "Artículos"
	case TablePrices: return "Precios"
	case TableReceipts: return "Recibos"
	case TableReceiptArticles: return "Artículos de los recibos"

	case TableStatusDate: return "Fecha de los estatus"
	}

	return ""
}
