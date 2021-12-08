package models

import (

	"github.com/go-ozzo/ozzo-validation"
)

func(v *sysValidation) CreateReceipt(model *Receipt) (err error){
	validations := []*validation.FieldRules{
		validation.Field(&model.AmountTotal,validation.Min(0.01)),
		validation.Field(&model.Code,validation.Required),
		validation.Field(&model.TableID, validation.Required),
		validation.Field(&model.ClientID, validation.Required),
		validation.Field(&model.AmountTotal, validation.Min(0.01)),
		validation.Field(&model.StateID, validation.Required),
		validation.Field(&model.CurrencyID, validation.Required),
		validation.Field(&model.ReceiptArticles, validation.Required),
	}


	err = validation.ValidateStruct(model, validations...)

	return
}

func(v *sysValidation) SaveReceipt(model *Receipt) (err error){
	validations := []*validation.FieldRules{
		validation.Field(&model.ID, validation.Required),
		validation.Field(&model.AmountTotal, validation.Min(0.01)),
		validation.Field(&model.StateID, validation.Required),
		validation.Field(&model.CurrencyID, validation.Required),
		validation.Field(&model.ReceiptArticles, validation.Required),
	}


	err = validation.ValidateStruct(model, validations...)

	return
}

