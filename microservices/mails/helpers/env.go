package helpers 

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ENV struct{
	NAME_SERVICE string
	NAME_MICROSERVICE_EMAIL string

	MAIL_FROM string
	MAIL_CC	string
	
	SMTP_HOST string
	SMTP_PORT	int 
	SMTP_USER string
	SMTP_PASSWORD string
}

var (
	Env *ENV = &ENV{}
)

func(env *ENV) Load(path string) {
	godotenv.Load(path)

	env.NAME_SERVICE = os.Getenv("NAME_SERVICE")

	env.MAIL_FROM = os.Getenv("MAIL_FROM")
	env.MAIL_CC = os.Getenv("MAIL_CC")

	env.SMTP_HOST = os.Getenv("SMTP_HOST")
	env.SMTP_PORT,_ = strconv.Atoi(os.Getenv("SMTP_PORT"))
	env.SMTP_USER = os.Getenv("SMTP_USER")
	env.SMTP_PASSWORD = os.Getenv("SMTP_PASSWORD")
}
