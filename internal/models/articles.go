package models

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Article struct {
	BaseModel	
	Name	string	`gorm:"not null;type:varchar(50)"`
	cant	uint	`gorm:"not null;"`
	
	SellerID	uuid.UUID	`gorm:"not null;type:uuid"`
	StateID	uint	`gorm:"not null"`
	TableID	uint	`gorm:"not null"`
	
	Seller	User	`gorm:"foreignKey:SellerID"`
	Table	ITable	`gorm:"foreignKey:TableID"`
	State	IState	`gorm:"foreignKey:StateID"`
}


func (model *Article) getFromDB(db *gorm.DB) *gorm.DB{
	return db.Model(model)
}
