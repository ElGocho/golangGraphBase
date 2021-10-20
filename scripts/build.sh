#!/bin/bash
#-*- ENCODING: UTF-8 -*-

echo "Compilando"
go build -o builds/sa_service main/main.go 2>> logs/buildLog.txt

exit
