package models

import (
	"gorm.io/gorm"
)

type ICountry struct {
	ID uint `gorm:"primaryKey;->"`
	Name	string	`gorm:"->;not null;type:varchar(50)"`
	Code	string	`gorm:"->;not null;unique;type:varchar(25)"`
	Status	bool	`gorm:"->;not null;default:true"`
}

type WCountry struct {
	ID uint `gorm:"primaryKey;<-"`
	Name	string	`gorm:"<-;not null;type:varchar(50)"`
	Code	string	`gorm:"<-;not null;unique;type:varchar(25)"`
	Status	bool	`gorm:"<-;not null;default:true"`

	Currencies	[]*WCurrency	`gorm:"foreignKey:CountryID"`
	Languages []*WLanguage	`gorm:"foreignKey:CountryID"`
}

func (models *WCountry) TableName() string{
	return "i_countries"
}

func (model *ICountry) getFromDB(db *gorm.DB) *gorm.DB{
	return db.Model(model)
}
