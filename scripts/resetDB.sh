#!/bin/bash
#-*- ENCODING: UTF-8 -*-

usuario=usuario
password=secret_password
host=localhost
port=54322
bdToDrop=miBD
bd=postgres

if [ -n "$1" ]; then 
	usuario=$1
fi

if [ -n "$2" ]; then
	bdToDrop=$2
fi

if [ -n "$3" ]; then
	password=$3
fi

if [ -n "$4" ]; then
	port=$4
fi

if [ -n "$5" ]; then
	host=$5
fi

echo "Reseteando DB"

PGPASSWORD=$password psql -U $usuario -h $host -p $port -d $bd -c "drop database \"$bdToDrop\""


PGPASSWORD=$password psql -U $usuario -h $host -p $port -d $bd -c "create database \"$bdToDrop\""
