package models

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
	"github.com/go-ozzo/ozzo-validation"

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

	Client	User	`gorm:"foreignKey:ClientID;->"`
	Table	ITable	`gorm:"foreignKey:TableID;->"`
	State	IState	`gorm:"foreignKey:StateID;->"`
	Currency	ICurrency	`gorm:"foreignKey:CurrencyID;->"`

	ReceiptArticles	[]*ReceiptArticle	`gorm:"foreignKey:ReceiptID"`
}


type Receipts	[]*Receipt


func (model *Receipt) BeforeCreate(tx *gorm.DB) (err error){
	model.ID = uuid.New()

	err = Repositories.ReceiptAmountTotal(tx,model)

	if err != nil {
		return
	}

	err = Repositories.ReceiptCode(tx,model)

	if err != nil {
		return 
	}

	validations := []*validation.FieldRules{
		validation.Field(&model.AmountTotal,validation.Min(0.01)),
		validation.Field(&model.Code,validation.Required),
		validation.Field(&model.TableID, validation.Required),
		validation.Field(&model.ClientID, validation.Required),
	}

	err = model.Validate(validations...)

	if err != nil {
		return 
	}

	return 
}

func (model *Receipt) BeforeSave(tx *gorm.DB) (err error){
	validations := []*validation.FieldRules{
		validation.Field(&model.ID, validation.Required),
	}
	
	err = model.Validate(validations...)

	if err != nil {
		return 
	}

	return 
}

func (model *Receipt) Get(tx *gorm.DB, builder *Builder) *gorm.DB{
	db := BuilderORMQuery(tx, builder)

	return db.First(model)
}

func (model *Receipts) Find(tx *gorm.DB, builder *Builder) *gorm.DB{
	db := BuilderORMQuery(tx, builder)

	return db.Find(&model)
}

func (model *Receipt) Create(db *gorm.DB) *gorm.DB{
	return db.Create(model)
}

func (model *Receipt) Save(tx *gorm.DB, builder *Builder) error{
	db := BuilderORMQuery(tx, builder)

	err := db.Save(model).Error

	if err != nil {
		return err
	}

	err = ReceiptArticles(model.ReceiptArticles).Save(tx, builder)

	return err
}

func (model Receipts) Save(tx *gorm.DB, builder *Builder) error{
	
	for _, receipt := range model {
		if err := receipt.Save(tx, builder); err != nil {
			return err	
		} 
	}

	return nil
}

func (model *Receipts) Create(db *gorm.DB) *gorm.DB{
	return db.Create(model)
}

func (model *Receipt) Validate(validations ...*validation.FieldRules) (err error){

	validations = append(validations, 
		validation.Field(&model.AmountTotal, validation.Min(0.01)),
		validation.Field(&model.StateID, validation.Required),
		validation.Field(&model.CurrencyID, validation.Required),
		validation.Field(&model.ReceiptArticles, validation.Required),
	)

	err = validation.ValidateStruct(model, validations...)

	return err
}
