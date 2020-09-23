#!/bin/bash

cd
#GOOS=linux go build main.go -ldflags "-X main.MODE PROD"
GOOS=linux go build -ldflags="-X 'github.com/ainsleyclark/verbis/api.SuperAdminString=false'"

rm -rf build/theme
rm -rf build/admin
rm -rf build/api
rm -rf build/verbis
rm -rf build/config

cp main build

mkdir build/api
mkdir build/api/database

cp -a api/database/migrations build/api/database
cp -a theme build/theme
cp -a admin build/admin
cp -a config build/config

mv build/main build/verbis