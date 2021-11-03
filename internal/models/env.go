package models

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ENV struct{
	ENVIRONMENT string
	AUTO_MIGRATE bool	
	NAME_SERVICE string

	DB_USER	string
	DB_PASSWORD string
	DB_NAME string
	DB_HOST string
	DB_PORT	string
	DB_SSLMODE string 

	GIN_MODE string
	PORT string

	LOG_MODE bool
}


func(env *ENV) Load(path string) {
	godotenv.Load(path)

	env.ENVIRONMENT = os.Getenv("ENVIRONMENT")
	env.AUTO_MIGRATE,_ = strconv.ParseBool(os.Getenv("AUTO_MIGRATE"))
	env.NAME_SERVICE = os.Getenv("NAME_SERVICE")
	env.DB_USER = os.Getenv("DB_USER")
	env.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	env.DB_NAME = os.Getenv("DB_NAME")
	env.DB_HOST = os.Getenv("DB_HOST")
	env.DB_PORT = os.Getenv("DB_PORT")
	env.DB_SSLMODE = os.Getenv("DB_SSLMODE")

	env.GIN_MODE = os.Getenv("GIN_MODE")
	env.PORT = os.Getenv("PORT")

	env.LOG_MODE,_ = strconv.ParseBool(os.Getenv("LOG_MODE"))

}
