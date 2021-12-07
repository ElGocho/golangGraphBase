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

func Welcome(db *gorm.DB, message helpers.Message) (err error){
	data,OK := message.Data.(string)

	if !OK {
		return fmt.Errorf("1-Data no valida, debe ser un string")
	}

	user := &models.User{}
	builder := &models.Builder{}

	where := models.Where{
		Condition: "id = ?",
		Params: []interface{}{data},
	}

	builder.Where = append(builder.Where, where)

	builder.Select = []string{"name", "email"}

	err = user.Get(db, builder).Error

	if err != nil {
		return 
	}

	templateData := struct {
		NAME_SERVICE	string
		UserName	string
	}{
		mailHelpers.Env.NAME_SERVICE,
		user.Name,
	}

	body, err := mailHelpers.LoadTemplate(consts.TemplateMailWelcome, templateData)

	if err != nil {
		return
	}

	err = mailHelpers.SendMail(user.Email, body,nil)

	return
}
