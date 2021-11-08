package models

import (
	"gorm.io/gorm"
	
	"sa_web_service/internal/models/consts"
)

type IRole struct{
	ID uint `gorm:"primaryKey;->"`
	Name	string	`gorm:"->;not null;type:varchar(50)"`
	Code	string	`gorm:"->;unique;not null;type:varchar(50)"`
	Status	bool	`gorm:"->;not null;default:true"`
}

type WRole struct{
	ID uint `gorm:"primaryKey;<-"`
	Name	string	`gorm:"<-;not null;type:varchar(50)"`
	Code	string	`gorm:"<-;unique;not null;type:varchar(50)"`
	Status	bool	`gorm:"<-;not null;default:true"`
}

func (model *WRole) TableName() string{
	return string(consts.TableRoles)
}

func (model *IRole) getFromDB(db *gorm.DB) *gorm.DB{
	return db.Model(model)
}
