package models

import (
	_"fmt"
	"gorm.io/gorm"
	"github.com/google/uuid"

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
	
	err = Validations.CreateArticle(model)

	if err != nil {
		return 
	}

	return 
}

func (model *Article) BeforeSave(tx *gorm.DB) (err error){
	err = Validations.SaveArticle(model)

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


