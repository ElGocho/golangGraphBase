#!/bin/bash
#-*- ENCODING: UTF-8 -*-

echo "Generando"
CGO_ENABLED=0 go get github.com/99designs/gqlgen
CGO_ENABLED=0 go run github.com/99designs/gqlgen

exit
