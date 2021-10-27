package models

import (
	"gorm.io/gorm"
)


type IState struct{
	ID uint `gorm:"primaryKey;->"`
	Name	string	`gorm:"->;not null;type:varchar(50)"`
	Code	string	`gorm:"->;unique;not null;type:varchar(25)"`
	Status	bool	`gorm:"->;not null;default:true"`
}

func (model *IState) getFromDB(db *gorm.DB) *gorm.DB{
	return db.Model(model)
}
