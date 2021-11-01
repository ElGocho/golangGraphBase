package models

import (
	"gorm.io/gorm"
	"github.com/google/uuid"

	"sa_web_service/internal/helpers"
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

	State	IState `gorm:"foreignKey:StateID;->"`
	Role	IRole	`gorm:"foreignKey:RoleID;->"`
	Table	ITable	`gorm:"foreignKey:TableID;->"`
}


func (model *User) BeforeCreate(tx *gorm.DB) (err error){
	model.ID = uuid.New()
	
	if model.Password != "" {
		hash, err := helpers.Encrypt(model.Password)

		if err != nil {
			return err
		}

		model.Password = hash
	}

	return 
}

func (model *User) Get(db *gorm.DB, builder *Builder) *gorm.DB{
	BuilderORMQuery(db, builder)

	return db.First(model)
}

func (model *User) Create(db *gorm.DB) *gorm.DB{
	return db.Create(model)
}


