package models

import (
	"gorm.io/gorm"

	"sa_web_service/internal/models/const"
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
	return string(cons.TableCountries)
}

func (model *ICountry) getFromDB(db *gorm.DB) *gorm.DB{
	return db.Model(model)
}
