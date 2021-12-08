package models

import (

	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func(v *sysValidation) CreateUser(model *User) (err error){
	validations := []*validation.FieldRules{
		validation.Field(&model.Name, validation.Required, validation.Length(3,50)),
		validation.Field(&model.Email, validation.Required, is.Email),
		validation.Field(&model.Password, validation.Required, is.Alphanumeric),
		validation.Field(&model.Identification, validation.Required, validation.Length(3,50)),
		validation.Field(&model.Image, validation.Required, is.Alphanumeric),
		validation.Field(&model.StateID, validation.Required),
		validation.Field(&model.RoleID, validation.Required),
		validation.Field(&model.TableID, validation.Required),
	}


	err = validation.ValidateStruct(model, validations...)

	return
}

func(v *sysValidation) SaveUser(model *User) (err error){
	validations := []*validation.FieldRules{}


	err = validation.ValidateStruct(model, validations...)

	return
}

