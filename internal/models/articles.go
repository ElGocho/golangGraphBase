package models

import (
	_"fmt"
	"gorm.io/gorm"
	"github.com/google/uuid"
	"github.com/go-ozzo/ozzo-validation"

)

type Article struct {
	BaseModel	
	Name	string	`gorm:"not null;type:varchar(50)"`
	Cant	uint	`gorm:"not null;"`
	
	SellerID	uuid.UUID	`gorm:"not null;type:uuid"`
	StateID	uint	`gorm:"not null"`
	TableID	uint	`gorm:"not null"`
	
	Seller	User	`gorm:"foreignKey:SellerID;->"`
	Table	ITable	`gorm:"foreignKey:TableID;->"`
	State	IState	`gorm:"foreignKey:StateID;->"`

	Prices []*Price	`gorm:"foreignKey:TableItemID"`
}

type Articles []*Article


func (model *Article) BeforeCreate(tx *gorm.DB) (err error){
	model.ID = uuid.New()
	
	validations := []*validation.FieldRules{
		validation.Field(&model.TableID, validation.Required),
		validation.Field(&model.SellerID, validation.Required),
	}

	err = model.Validate(validations...)

	if err != nil {
		return 
	}

	return 
}

func (model *Article) BeforeSave(tx *gorm.DB) (err error){
	validations := []*validation.FieldRules{
		validation.Field(&model.ID, validation.Required),
	}
	
	err = model.Validate(validations...)

	if err != nil {
		return 
	}

	return 
}

func (model *Article) Get(tx *gorm.DB, builder *Builder) *gorm.DB{
	db := BuilderORMQuery(tx, builder)

	return db.First(model)
}

func (model *Articles) Find(tx *gorm.DB, builder *Builder) *gorm.DB{
	db := BuilderORMQuery(tx, builder)

	return db.Find(&model)
}

func (model *Article) Create(db *gorm.DB) *gorm.DB{
	return db.Create(model)
}

func (model *Article) Save(tx *gorm.DB, builder *Builder) error{
	db := BuilderORMQuery(tx, builder)

	err := db.Save(model).Error

	if err != nil {
		return err
	}

	err = Prices(model.Prices).Save(tx, builder)

	return err
}

func (model Articles) Save(tx *gorm.DB, builder *Builder) error{
	
	for _, article := range model {
		if err := article.Save(tx, builder); err != nil {
			return err	
		} 
	}

	return nil
}

func (model *Articles) Create(db *gorm.DB) *gorm.DB{
	return db.Create(model)
}

func (model *Article) Validate(validations ...*validation.FieldRules) (err error){

	var minCant uint = 1

	validations = append(validations, 
		validation.Field(&model.Name, validation.Required, validation.Length(3,50)),
		validation.Field(&model.Cant, validation.Required, validation.Min(minCant)),
		validation.Field(&model.StateID, validation.Required),
		validation.Field(&model.Prices, validation.Required),
	)

	err = validation.ValidateStruct(model, validations...)

	return err
}

