package models

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
	"time"
)


type Base struct {
	ID	uuid.UUID	`gorm:"type:uuid;primaryKey"`
}

type BaseModel struct{
	Base
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	gorm.DeletedAt	`gorm:"index"`
}

type BaseModelNoDeleted struct{
	Base
	CreatedAt	time.Time
	UpdatedAt	time.Time
}


func (base *Base) BeforeCreate(tx *gorm.DB) (err error){
	base.ID = uuid.New()

	return
}
