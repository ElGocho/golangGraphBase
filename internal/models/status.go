package models

import (
	"gorm.io/gorm"
)


type IState struct{
	ID uint `gorm:"primaryKey;->"`
	Name	string	`gorm:"->;not null;type:varchar(50)"`
	Code	string	`gorm:"->;unique;not null;type:varchar(25)"`
	Status	bool	`gorm:"->;not null;default:true"`

	TableID	uint	`gorm:"->;not null;"`

	Table	ITable	`gorm:"->"`
}

type WState struct{
	ID uint `gorm:"primaryKey;<-"`
	Name	string	`gorm:"<-;not null;type:varchar(50)"`
	Code	string	`gorm:"<-;unique;not null;type:varchar(25)"`
	Status	bool	`gorm:"<-;not null;default:true"`

	TableID	uint	`gorm:"<-;not null;"`
}

func (model *WState) TableName() string {
	return "i_states"
}

func (model *IState) Get(db *gorm.DB, builder *Builder) *gorm.DB{
	BuilderORMQuery(db, builder)

	return db.First(model)
}

