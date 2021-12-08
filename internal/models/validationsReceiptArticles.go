package models

import (

	"github.com/go-ozzo/ozzo-validation"
)

func(v *sysValidation) CreateReceiptArticle(model *ReceiptArticle) (err error){
	validations := []*validation.FieldRules{
		validation.Field(&model.ReceiptID, validation.Required),
		validation.Field(&model.ArticleID, validation.Required),
		validation.Field(&model.TableID, validation.Required),
		validation.Field(&model.PriceID, validation.Required),
		validation.Field(&model.StateID, validation.Required),
	}


	err = validation.ValidateStruct(model, validations...)

	return
}

func(v *sysValidation) SaveReceiptArticle(model *ReceiptArticle) (err error){
	validations := []*validation.FieldRules{
		validation.Field(&model.ID, validation.Required),
		validation.Field(&model.PriceID, validation.Required),
		validation.Field(&model.StateID, validation.Required),
	}


	err = validation.ValidateStruct(model, validations...)

	return
}
