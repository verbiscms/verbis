#!/bin/bash

# Remove and create the build directory
rm -rf ./build
mkdir ./build

function build() {

	# Get the build path
	path=./build/$1
	mkdir $path

    if [[ $1 == "mac" ]]
		then
			echo "Building for mac"
			CGO_ENABLED=1 go build -o "verbis" -ldflags="-X 'github.com/ainsleyclark/verbis/api.SuperAdminString=false'"
		else
			echo "Building for $1"
		  	CGO_ENABLED=1 GOOS=$1 go build -o "verbis" -ldflags="-X 'github.com/ainsleyclark/verbis/api.SuperAdminString=false'"
    fi

	# API (Database)
	mkdir $path/api/
	mkdir $path/api/database && cp -a api/database/migrations/schema.sql $path/api/database

	# Theme
	mkdir $path/theme && cp -a theme $path/theme

	# Admin
	mkdir $path/admin && cp -a admin $path/admin

	# Config
	mkdir $path/config && cp -a config $path/config

	## TODO: Add storage paths
}

# Build for
build mac
build linux
build windows