package helpers

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm/logger"

)

func InitDB(user, password, host, port, dbName, sslMode string, logMode bool) (db *gorm.DB,err error){
	myLog := logger.Error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbName, port, sslMode)

	if logMode{
		myLog = logger.Info
	}

	db,err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(myLog),
	})

	return db, err
}

func NewSession(tx *gorm.DB) *gorm.DB{
	return tx.Session(&gorm.Session{NewDB: true})
}
