package helpers 

import (
	"bytes"
	"html/template"

	"sa_web_service/internal/models/consts"
)


func LoadTemplate(src consts.Template, param interface{}) (resp string, err error){
	t := template.New(src.Name())

	t, err = t.ParseFiles(string(src))

	if err != nil {
		return
	}

	var tpl bytes.Buffer

	err = t.Execute(&tpl, param)

	if err != nil {
		return
	}

	resp = tpl.String()

	return
}
