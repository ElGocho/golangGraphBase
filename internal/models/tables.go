package models

import (
	"gorm.io/gorm"
)


type ITable struct{
	ID uint `gorm:"primaryKey;->"`
	Name	string	`gorm:"->;not null;type:varchar(50)"`
	Code	string	`gorm:"->;unique;not null;type:varchar(50)"`
	Status	bool	`gorm:"->;not null;default:true"`

	Tables []interface{} `gorm:"-"`
}

type WTable struct{
	ID uint `gorm:"primaryKey;<-"`
	Name	string	`gorm:"<-;not null;type:varchar(50)"`
	Code	string	`gorm:"<-;unique;not null;type:varchar(50)"`
	Status	bool	`gorm:"<-;not null;default:true"`
}

func (model *WTable) TableName() string{
	return "i_tables"
}

func (model *ITable) getFromDB(db *gorm.DB) *gorm.DB{
	return db.Model(model)
}
