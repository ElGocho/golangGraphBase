package models

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ENV struct{
	ENVIRONMENT string
	AUTO_MIGRATE bool	
	PORT string
	NAME_SERVICE string

	DB_USER	string
	DB_PASSWORD string
	DB_NAME string
	DB_HOST string
	DB_PORT	string
	DB_SSLMODE string 
}


func(env *ENV) Load(path string) {
	godotenv.Load(path)

	env.ENVIRONMENT = os.Getenv("ENVIRONMENT")
	env.AUTO_MIGRATE,_ = strconv.ParseBool(os.Getenv("AUTO_MIGRATE"))
	env.PORT = os.Getenv("PORT")
	env.NAME_SERVICE = os.Getenv("NAME_SERVICE")
	env.DB_USER = os.Getenv("DB_USER")
	env.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	env.DB_NAME = os.Getenv("DB_NAME")
	env.DB_HOST = os.Getenv("DB_HOST")
	env.DB_PORT = os.Getenv("DB_PORT")
	env.DB_SSLMODE = os.Getenv("DB_SSLMODE")
}
