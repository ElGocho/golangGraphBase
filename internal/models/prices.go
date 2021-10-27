package models

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Price struct {
	BaseModel	
	Amount	float64 `gorm:"not null;type:numeric(18,2)"`
	Status	bool	`gorm:"not null;"`
	
	TableItemID	uuid.UUID	`gorm:"not null;type:uuid"`
	CurrencyID	uint	`gorm:"not null;"`
	FTableID	uint	`gorm:"not null"`
	
	Table	ITable `gorm:"foreignKey:FTableID"`
	Currency ICurrency	`gorm:"foreignKey:CurrencyID"` 

	TableItems	interface{}	`gorm:"-"`
}


func (model *Price) getFromDB(db *gorm.DB) *gorm.DB{
	return db.Model(model)
}
