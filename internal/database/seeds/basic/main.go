package basic


func GetAll() (resp []interface{}){

	resp = append(resp, GetWTables())
	resp = append(resp, GetWCountries())
	resp = append(resp, GetWRoles())

	return
}


