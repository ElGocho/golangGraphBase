package models

import (

	"github.com/go-ozzo/ozzo-validation"
)

func(v *sysValidation) CreateArticle(model *Article) (err error){
	minCant := 1

	validations := []*validation.FieldRules{
		validation.Field(&model.TableID, validation.Required),
		validation.Field(&model.SellerID, validation.Required),
		validation.Field(&model.Name, validation.Required, validation.Length(3,50)),
		validation.Field(&model.Cant, validation.Required, validation.Min(minCant)),
		validation.Field(&model.StateID, validation.Required),
		validation.Field(&model.Prices, validation.Required),
	}


	err = validation.ValidateStruct(model, validations...)

	return
}

func(v *sysValidation) SaveArticle(model *Article) (err error){
	minCant := 1
	
	validations := []*validation.FieldRules{
		validation.Field(&model.ID, validation.Required),
		validation.Field(&model.Name, validation.Required, validation.Length(3,50)),
		validation.Field(&model.Cant, validation.Required, validation.Min(minCant)),
		validation.Field(&model.StateID, validation.Required),
		validation.Field(&model.Prices, validation.Required),
	}


	err = validation.ValidateStruct(model, validations...)

	return
}


