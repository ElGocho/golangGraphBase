package models

import (
	"gorm.io/gorm"

	"sa_web_service/internal/models/const"
)

type ICurrency struct {
	ID uint `gorm:"primaryKey;->"`
	Name	string	`gorm:"->;not null;type:varchar(50)"`
	Code	string	`gorm:"->;not null;unique;type:varchar(25)"`
	Status	bool	`gorm:"->;not null;default:true"`
	CountryID	uint	`gorm:"->;not null;"`

	Country	ICountry	`gorm:"foreignKey:CountryID"`
}

type WCurrency struct {
	ID uint `gorm:"primaryKey;<-"`
	Name	string	`gorm:"<-;not null;type:varchar(50)"`
	Code	string	`gorm:"<-;not null;unique;type:varchar(25)"`
	Status	bool	`gorm:"<-;not null;default:true"`
	CountryID	uint	`gorm:"<-;not null;"`
}

func (model *WCurrency) TableName() string{
	return string(cons.TableCurrencies)
}

func (model *ICurrency) getFromDB(db *gorm.DB) *gorm.DB{
	return db.Model(model)
}
