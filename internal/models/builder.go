package models

import (
	"gorm.io/gorm"
)

type Builder struct {
	Table	string	`json:"table"`
	Select	interface{} `json:"select"`
	Where	[]interface{}	`json:"where"`
}


func NewSession(tx *gorm.DB) *gorm.DB{
	return tx.Session(&gorm.Session{NewDB: true})
}

func BuilderORMQuery(db *gorm.DB, builder *Builder){
	if(builder == nil){
		return 
	}

	BuilderORMTable(db,builder)

	BuilderORMSelect(db, builder)

}

func BuilderORMTable(db *gorm.DB, builder *Builder){
	if(builder == nil || builder.Table == ""){
		return
	}

	db = db.Table(builder.Table)
}

func BuilderORMSelect(db *gorm.DB, builder *Builder) {
	if(builder == nil || builder.Select == nil){
		return
	}

	db = db.Select(builder.Select)
}


func SelectOnlyID() *Builder{
	return &Builder{ Select: "id" }
}
