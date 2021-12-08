package models

import (
	_"fmt"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Price struct {
	BaseModel	
	Amount	float64 `gorm:"not null;type:numeric(18,2)"`
	Status	bool	`gorm:"not null;default:true"`
	
	TableItemID	uuid.UUID	`gorm:"not null;type:uuid"`
	CurrencyID	uint	`gorm:"not null;"`
	FTableID	uint	`gorm:"not null"`
	
	FTable	ITable `gorm:"foreignKey:FTableID;->"`
	Currency ICurrency	`gorm:"foreignKey:CurrencyID;->"` 

	TableItem	interface{}	`gorm:"-"`
}

type Prices []*Price

func (model *Price) BeforeCreate(tx *gorm.DB) (err error){
	model.ID = uuid.New()

	err = Validations.CreatePrice(model)

	if err != nil {
		return 
	}

	return 
}

func (model *Price) BeforeSave(tx *gorm.DB) (err error){
	err = Validations.SavePrice(model)

	if err != nil {
		return 
	}

	return 
}

func (model *Price) Get(tx *gorm.DB, builder *Builder) *gorm.DB{
	db := BuilderORMQuery(tx, builder)

	return db.First(model)
}

func (model *Prices) Find(tx *gorm.DB, builder *Builder) *gorm.DB{
	db := BuilderORMQuery(tx, builder)

	return db.Find(&model)
}

func (model *Price) Create(db *gorm.DB) *gorm.DB{
	return db.Create(model)
}

func (model *Price) Save(tx *gorm.DB, builder *Builder) error{
	db := BuilderORMQuery(tx, builder)

	return db.Save(model).Error
}

func (model Prices) Save(tx *gorm.DB, builder *Builder) error{
	for _, price := range model {
		if err := price.Save(tx, builder); err != nil {
			return err
		}
	}

	return nil
}


