#!/bin/bash

cd api
GOOS=linux go build main.go

cd ..
rm -rf build/theme
rm -rf build/admin
rm build/api/cms

cp api/main build/api/cms
cp -a api/database/migrations build/api/database/migrations
cp -a theme build/theme
cp -a admin build/admin

