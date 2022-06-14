#! /bin/bash
# Author: arvryna on Ср 06 апр 2022 09:50:24 MSK
# desc: Build and export to bin
echo "Building.."
go build -o $GOPATH/bin/sled main.go

