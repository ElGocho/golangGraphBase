package consts

//Enumeración de Tablas
const (
	TableTables Table = "i_tables"
	TableCountries Table = "i_countries"
	TableLanguages Table = "i_languages"
	TableStates Table = "i_states"
	TableRoles	Table =	"i_roles"
	TableCurrencies Table = "i_currencies"
	TableUsers Table = "users"
	TableArticles Table = "articles"
	TablePrices Table = "prices"
	TableReceipts Table = "receipts"
	TableReceiptArticles Table = "receipt_articles"

	TableStatusDate Table = "status_date"
)

func (t Table) Name() string{
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
