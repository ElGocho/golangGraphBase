#!/bin/bash
#-*- ENCODING: UTF-8 -*-

echo "Compilando"
go build -o builds/mails_microservice microservices/mails/main/main.go 2>> logs/buildMailsMicroserviceLog.txt

exit

