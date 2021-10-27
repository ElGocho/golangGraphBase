package models

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
	"time"
)

type StatusDate struct {
	BaseModel
	Date	time.Time	`gorm:"not null"`

	TableItemID	uuid.UUID	`gorm:"not null;type:uuid"`
	FTableID	uint	`gorm:"not null;"`
	StateID	uint	`gorm:"not null"`

	Table	ITable	`gorm:"foreignKey:FTableID"`	
	State	IState	`gorm:"foreignKey:StateID"`	

	TableItems	[]interface{}	`gorm:"-"`
}


func (model *StatusDate) getFromDB(db *gorm.DB) *gorm.DB{
	return db.Model(model)
}
