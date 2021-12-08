package models

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type ReceiptArticle struct {
	BaseModel
	
	ReceiptID	uuid.UUID	`gorm:"not null;type:uuid"`
	ArticleID	uuid.UUID	`gorm:"not null;type:uuid"`
	PriceID	uuid.UUID	`gorm:"not null;type:uuid"`
	TableID	uint	`gorm:"not null"`
	StateID	uint	`gorm:"not null"`

	Receipt	Receipt	`gorm:"foreignKey:ReceiptID;->"`
	Article	Article	`gorm:"foreignKey:ArticleID;->"`
	Price	Price		`gorm:"foreignKey:PriceID;->"`
	Table	ITable	`gorm:"foreignKey:TableID;->"`
	State	IState	`gorm:"foreignKey:StateID;->"`
}


type ReceiptArticles	[]*ReceiptArticle


func (model *ReceiptArticle) BeforeCreate(tx *gorm.DB) (err error){
	model.ID = uuid.New()
	
	err = Validations.CreateReceiptArticle(model)

	if err != nil {
		return 
	}

	return 
}

func (model *ReceiptArticle) BeforeSave(tx *gorm.DB) (err error){
	err = Validations.SaveReceiptArticle(model)

	if err != nil {
		return 
	}

	return 
}

func (model *ReceiptArticle) Get(tx *gorm.DB, builder *Builder) *gorm.DB{
	db := BuilderORMQuery(tx, builder)

	return db.First(model)
}

func (model *ReceiptArticles) Find(tx *gorm.DB, builder *Builder) *gorm.DB{
	db := BuilderORMQuery(tx, builder)

	return db.Find(&model)
}

func (model *ReceiptArticle) Create(db *gorm.DB) *gorm.DB{
	return db.Create(model)
}

func (model *ReceiptArticle) Save(tx *gorm.DB, builder *Builder) error{
	db := BuilderORMQuery(tx, builder)

	err := db.Save(model).Error

	return err
}

func (model ReceiptArticles) Save(tx *gorm.DB, builder *Builder) error{
	
	for _, receipt := range model {
		if err := receipt.Save(tx, builder); err != nil {
			return err	
		} 
	}

	return nil
}

func (model *ReceiptArticles) Create(db *gorm.DB) *gorm.DB{
	return db.Create(model)
}

