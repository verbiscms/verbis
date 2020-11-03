#!/bin/bash

# Remove and create the build directory
rm -rf ./build/mac
rm -rf ./build/linux
rm -rf ./build/windows

function build() {

	# Get the build path
	os=$1
	path=./build/$1
	mkdir $path

	# Mac
    if [[ $1 == "mac" ]]
		then
			echo "Building for mac..."
			CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o "./build/mac/verbis" -ldflags="-X 'github.com/ainsleyclark/verbis/api.SuperAdminString=false'"
    fi

	# Linux
    if [[ $1 == "linux" ]]
		then
			echo "Building for linux..."
		  	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "./build/linux/verbis" -ldflags="-X 'github.com/ainsleyclark/verbis/api.SuperAdminString=false'"
    fi

	# Windows
    if [[ $1 == "windows" ]]
		then
			echo "Building for windows..."
		  	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o "/build/windows/verbis.exe" -ldflags="-X 'github.com/ainsleyclark/verbis/api.SuperAdminString=false'"
    fi

	# Build Vue
	cd admin
	npm run --silent --no-progress build
	cd ../

	# Build Web
	cd api/webp
	npm run --silent --no-progress prod
	cd ../../

	# Env
	cp .env.example $path/.env.example

	# API (Database)
	mkdir $path/api/
	mkdir $path/api/database && cp -a api/database/migrations/schema.sql $path/api/database

	# API (Web)
	mkdir $path/api/web && rsync -av --quiet api/web/ build/$os/api/web/ --exclude node_modules

	# Theme
	mkdir $path/theme && cp -a theme/ $path/theme/

	# Admin
	mkdir $path/admin
	rsync -av --quiet admin/dist/ build/$os/admin/

	# Config
	mkdir $path/config/ && cp -a config/ $path/config/

	# Storage
	mkdir $path/storage && touch .keep $path/storage
	mkdir $path/storage/fields && touch .keep $path/storage/fields
	mkdir $path/storage/uploads && touch .keep $path/storage/uploads
	mkdir $path/storage/dumps && touch .keep $path/storage/dumps
	mkdir $path/storage/logs && touch .keep $path/storage/logs

	# .gitignore
	printf 'node_modules\n.env\n.env.local\n.env.*.local\n.idea\n.vscode\n*.suo\n*.ntvs*\n*.njsproj\n*.sln\n*.sw?\n.DS_Store\n' > $path/.gitignore

	#keep

	echo "Build for "$os" completed"
}

# Build for
build mac
build linux
build windows