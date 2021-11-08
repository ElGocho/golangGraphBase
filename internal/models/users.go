package models

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"

	"sa_web_service/internal/helpers"
)

type User struct {
	BaseModel
	Name	string	`gorm:"not null;type:varchar(50)"`
	Email	string	`gorm:"not null;unique;type:varchar(50)"`
	Identification string	`gorm:"not null;type:varchar(50)"`
	Image	string	`gorm:"not null;type:varchar(75)"`
	Password	string	`gorm:"not null;type:varchar(255)"`
	
	TableID	uint	`gorm:"not null"`
	StateID	uint	`gorm:"not null"`
	RoleID	uint	`gorm:"not null"`

	Table	ITable	`gorm:"foreignKey:TableID;->"`
	State	IState `gorm:"foreignKey:StateID,TableID;references:ID,TableID;->"`
	Role	IRole	`gorm:"foreignKey:RoleID;->"`
}

type Users []*User

func (model *User) BeforeCreate(tx *gorm.DB) (err error){
	model.ID = uuid.New()
	
	err = model.Validate()

	if err != nil {
		return 
	}

	if model.Password != "" {
		hash, err := helpers.Hash(model.Password)

		if err != nil {
			return err
		}

		model.Password = hash
	}

	return 
}

func (model *User) Get(tx *gorm.DB, builder *Builder) *gorm.DB{
	db := BuilderORMQuery(tx, builder)

	return db.First(model)
}

func (model *Users) Find(tx *gorm.DB, builder *Builder) *gorm.DB{
	db := BuilderORMQuery(tx, builder)

	return db.Find(&model)
}

func (model *User) Create(db *gorm.DB) *gorm.DB{
	return db.Create(model)
}

func (model *User) Validate() (err error){
	err = validation.ValidateStruct(model,
		validation.Field(&model.Name, validation.Required, validation.Length(3,50)),
		validation.Field(&model.Email, validation.Required, is.Email),
		validation.Field(&model.Password, validation.Required, is.Alphanumeric),
		validation.Field(&model.Identification, validation.Required, validation.Length(3,50)),
		validation.Field(&model.Image, validation.Required, is.Alphanumeric),
		validation.Field(&model.StateID, validation.Required),
		validation.Field(&model.RoleID, validation.Required),
		validation.Field(&model.TableID, validation.Required),
	)

	return err
}
