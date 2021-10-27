package models

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Receipt struct {
	BaseModel
	AmountTotal	float64	`gorm:"not null;type:numeric(18,2)"`
	AmountAccepted	float64	`gorm:"type:numeric(18,2)"`
	Code	string	`gorm:"not null;unique;type:varchar(255)"`
	
	ClientID	uuid.UUID	`gorm:"not null;type:uuid"`
	TableID	uint	`gorm:"not null"`
	StateID	uint	`gorm:"not null"`
	CurrencyID	uint	`gorm:"not null"`

	Client	User	`gorm:"foreignKey:ClientID;"`
	Table	ITable	`gorm:"foreignKey:TableID;"`
	State	IState	`gorm:"foreignKey:StateID;"`
	Currency	ICurrency	`gorm:"foreignKey:CurrencyID"`
}


func (model *Receipt) getFromDB(db *gorm.DB) *gorm.DB{
	return db.Model(model)
}
