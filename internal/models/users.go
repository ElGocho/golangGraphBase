package models

import (
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Name	string	`gorm:"not null;type:varchar(50)"`
	Email	string	`gorm:"not null;unique;type:varchar(50)"`
	Identification string	`gorm:"not null;type:varchar(50)"`
	Image	string	`gorm:"not null;type:varchar(75)"`
	Password	string	`gorm:"not null;type:varchar(255)"`
	
	StateID	uint	`gorm:"not null"`
	RoleID	uint	`gorm:"not null"`
	TableID	uint	`gorm:"not null"`

	State	IState `gorm:"foreignKey:StateID"`
	Role	IRole	`gorm:"foreignKey:RoleID"`
	Table	ITable	`gorm:"foreignKey:TableID"`
}


func (model *User) getFromDB(db *gorm.DB) *gorm.DB{
	return db.Model(model)
}
