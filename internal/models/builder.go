package models

import (
	"gorm.io/gorm"
)

type Builder struct {
	Table	string	`json:"table"`
	Select	interface{} `json:"select"`
	Where	[]Where	`json:"where"`
}

type Where struct {
	Condition interface{}
	Params []interface{}
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

	BuilderORMWhere(db, builder)
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

func BuilderORMWhere(db *gorm.DB, builder *Builder){
	if(builder == nil ){
		return
	}

	for _,w := range builder.Where {
		db = db.Where(w.Condition, w.Params...)	
	} 
}


func SelectOnlyID() *Builder{
	return &Builder{ Select: "id" }
}
