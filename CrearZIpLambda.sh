#!/bin/bash

# Seteamos solo para este script que genere un ejecutable para OS linux
export GOOS=linux

# Seteamos solo para este script que genere un ejecutable para arquitectura amd64
export GOARCH=amd64

# Generamos el ejecutable
# go build main.go
go build -tags lambda.norpc -o bootstrap main.go

# ******************************************************************************

# Amazon pide que las lambdas se suban zipeadas

# Elimina main.zip en caso de existir, en caso contrario nos indica que no se encuentra el archivo.
rm -f main.zip

# Creamos archivo zip (zipeamos)
zip main.zip bootstrap
