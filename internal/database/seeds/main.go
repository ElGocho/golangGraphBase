package seeds

import (
	"sa_web_service/internal/database/seeds/basic"
)

func GetAll() (resp []interface{}) {

	resp = basic.GetAll()

	return
}
