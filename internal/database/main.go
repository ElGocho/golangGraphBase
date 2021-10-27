package database

import (
	"gorm.io/gorm"

	"sa_web_service/internal/database/migrations"
	"sa_web_service/internal/database/seeds"
)


func ExecAll(db *gorm.DB){
	if(VerifyEmptyDB(db)){
		ExecMigrations(db)	
		ExecSeeds(db)
	}

}


func ExecMigrations(db *gorm.DB){
	structs := migrations.GetAll()

	db.AutoMigrate(structs...)
}


func ExecSeeds(db *gorm.DB){
	structs := seeds.GetAll()

	for _,s := range structs {
		db.Create(s)
	}

}

func VerifyEmptyDB(db *gorm.DB) bool{
	return !db.Migrator().HasTable("i_countries")
}


