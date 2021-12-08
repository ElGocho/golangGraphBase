package models

import (

	"github.com/go-ozzo/ozzo-validation"
)

func(v *sysValidation) CreatePrice(model *Price) (err error){
	validations := []*validation.FieldRules{
		validation.Field(&model.FTableID, validation.Required),
		validation.Field(&model.Amount, validation.Required, validation.Min(0.01)),
		validation.Field(&model.TableItemID, validation.Required),
		validation.Field(&model.CurrencyID, validation.Required),
	}


	err = validation.ValidateStruct(model, validations...)

	return
}

func(v *sysValidation) SavePrice(model *Price) (err error){
	validations := []*validation.FieldRules{
		validation.Field(&model.ID, validation.Required),
		validation.Field(&model.Amount, validation.Required, validation.Min(0.01)),
		validation.Field(&model.TableItemID, validation.Required),
		validation.Field(&model.CurrencyID, validation.Required),
	}


	err = validation.ValidateStruct(model, validations...)

	return
}



