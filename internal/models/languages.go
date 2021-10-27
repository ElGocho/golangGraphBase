package models

import (
	"gorm.io/gorm"
)

type ILanguage struct {
	ID uint `gorm:"primaryKey;->"`
	Name	string	`gorm:"->;not null;type:varchar(50)"`
	Code	string	`gorm:"->;not null;unique;type:varchar(50)"`
	Status	bool	`gorm:"->;not null;default:true"`
	CountryID	uint	`gorm:"->;not null;"`

	Country	ICountry `gorm:"foreignKey:CountryID"`
}

type WLanguage struct {
	ID uint `gorm:"primaryKey;<-"`
	Name	string	`gorm:"<-;not null;type:varchar(50)"`
	Code	string	`gorm:"<-;not null;unique;type:varchar(50)"`
	Status	bool	`gorm:"<-;not null;default:true"`
	CountryID	uint	`gorm:"<-;not null;"`
}

func (model *WLanguage) TableName() string {
	return "i_languages"
}

func (model *ILanguage) getFromDB(db *gorm.DB) *gorm.DB{
	return db.Model(model)
}