package events

import (
	"fmt"

	"gorm.io/gorm"
	_"github.com/google/uuid"

	"sa_web_service/internal/helpers"
	"sa_web_service/internal/models"
	"sa_web_service/internal/models/consts"

	mailHelpers "sa_web_service/microservices/mails/helpers"
)


func ReceiptsCreate(db *gorm.DB, message helpers.Message) (err error){
	data, OK := message.Data.([]interface{})

	if !OK {
		return fmt.Errorf("1-Data no valida, debe ser un []string")
	}

	users := &models.Users{}
	builder := &models.Builder{}

	builder.Select = []string{"email"}

	where := models.Where{
		Condition: "id IN (?)",
		Params: data,
	}

	builder.Where = append(builder.Where, where)

	err = users.Find(db, builder).Error

	if err != nil {
		return 
	}

	if len(*users) == 0{
		return fmt.Errorf("No se encontro usuarios") 
	}

	body, err := mailHelpers.LoadTemplate(consts.TemplateMailReceiptCreate, mailHelpers.Env)

	if err != nil {
		return
	}

	for _, user := range *users {
		err = mailHelpers.SendMail(user.Email, body,nil)

		if err != nil {
			return
		}
	}

	return
}
