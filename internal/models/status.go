package models

import (
	"gorm.io/gorm"
	
	"sa_web_service/internal/models/consts"
)


type IState struct{
	ID uint `gorm:"primaryKey;->;uniqueIndex:idx_id_table_id"`
	Name	string	`gorm:"->;not null;type:varchar(50)"`
	Code	string	`gorm:"->;unique;not null;type:varchar(25)"`
	Status	bool	`gorm:"->;not null;default:true"`

	TableID	uint	`gorm:"->;not null;uniqueIndex:idx_id_table_id"`

	Table	ITable	`gorm:"->"`
}

type WState struct{
	ID uint `gorm:"primaryKey;<-;uniqueIndex:idx_id_table_id"`
	Name	string	`gorm:"<-;not null;type:varchar(50)"`
	Code	string	`gorm:"<-;unique;not null;type:varchar(25)"`
	Status	bool	`gorm:"<-;not null;default:true"`

	TableID	uint	`gorm:"<-;not null;uniqueIndex:idx_id_table_id"`
}

func (model *WState) TableName() string {
	return string(consts.TableStates)
}

func (model *IState) Get(db *gorm.DB, builder *Builder) *gorm.DB{
	BuilderORMQuery(db, builder)

	return db.First(model)
}

