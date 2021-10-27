package models

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type ReceiptArticles struct {
	BaseModel
	
	SellerID	uuid.UUID	`gorm:"not null;type:uuid"`
	ReceiptID	uuid.UUID	`gorm:"not null;type:uuid"`
	ArticleID	uuid.UUID	`gorm:"not null;type:uuid"`
	PriceID	uuid.UUID	`gorm:"not null;type:uuid"`
	TableID	uint	`gorm:"not null"`
	StateID	uint	`gorm:"not null"`

	Seller	User	`gorm:"foreignKey:SellerID;"`
	Receipt	Receipt	
	Article	Article	
	Price	Price	
	Table	ITable	`gorm:"foreignKey:TableID;"`
	State	IState	`gorm:"foreignKey:StateID;"`
}


func (model *ReceiptArticles) getFromDB(db *gorm.DB) *gorm.DB{
	return db.Model(model)
}
