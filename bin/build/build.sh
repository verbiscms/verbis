#!/bin/bash
#
# build.sh
#
# How to use:
# Run ./bin/build/build.sh from the root to get the right paths.
# Make sure to pass through a commit message to push the
# changes to the TryVerbis repo, for example:
# ./bin/build/build.sh "Commit message"
#

# Set variables
commitmsg=$1

# Remove and create the build directory
rm -rf ./build/mac
rm -rf ./build/linux
rm -rf ./build/windows

function build() {

	if [[ $commitmsg == "" ]]
		then
			echo "Add commit message"
			exit
	fi

	# Get the build path
	os=$1
	path=./build/$1
	mkdir $path

	# Mac
    if [[ $1 == "mac" ]]
		then
			echo "Building for mac..."
			CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o "./build/mac/verbis" -ldflags="-X 'github.com/ainsleyclark/verbis/api.SuperAdminString=false' -X 'github.com/ainsleyclark/verbis/api.Stack=heythere'"
			CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "./build/mac/verbis-linux" -ldflags="-X 'github.com/ainsleyclark/verbis/api.SuperAdminString=false'"
			CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o "./build/mac/verbis-windows.exe" -ldflags="-X 'github.com/ainsleyclark/verbis/api.SuperAdminString=false'"
    fi

	# Linux
    if [[ $1 == "linux" ]]
		then
			echo "Building for linux..."
		  	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "./build/linux/verbis" -ldflags="-X 'github.com/ainsleyclark/verbis/api.SuperAdminString=false'"
		  	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o "./build/linux/verbis-mac" -ldflags="-X 'github.com/ainsleyclark/verbis/api.SuperAdminString=false'"
		  	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o "./build/linux/verbis-windows.exe" -ldflags="-X 'github.com/ainsleyclark/verbis/api.SuperAdminString=false'"
    fi

	# Windows
    if [[ $1 == "windows" ]]
		then
			echo "Building for windows..."
		  	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o "./build/windows/verbis.exe" -ldflags="-X 'github.com/ainsleyclark/verbis/api.SuperAdminString=false'"
		  	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o "./build/windows/verbis-mac" -ldflags="-X 'github.com/ainsleyclark/verbis/api.SuperAdminString=false'"
		  	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "./build/windows/verbis-linux" -ldflags="-X 'github.com/ainsleyclark/verbis/api.SuperAdminString=false'"
    fi

	# Build Vue
	cd admin
	npm run --silent --no-progress build
	cd ../

	# Build Web
	cd api/web
	npm run --silent --no-progress prod
	cd ../../

	# Env
	cp .env.example $path/.env.example

	# API (Database)
	mkdir $path/api/
	mkdir $path/api/database && cp -a api/database/migrations/schema.sql $path/api/database

	# API (Web)
	mkdir $path/api/web && rsync -av --quiet api/web/ build/$os/api/web/ --exclude node_modules --exclude src --exclude package.json --exclude package-lock.json

	# Theme
	mkdir $path/theme && cp -a theme/ $path/theme/

	# Admin
	mkdir $path/admin
	rsync -av --quiet admin/dist/ build/$os/admin/

	# Config
	mkdir $path/config/ && cp -a config/ $path/config/

	# Storage
	mkdir $path/storage && touch $path/storage/.keep
	mkdir $path/storage/fields && touch $path/storage/fields/.keep
	mkdir $path/storage/uploads
	mkdir $path/storage/dumps && touch $path/storage/dumps/.keep
	mkdir $path/storage/logs && touch $path/storage/logs/.keep

	# Mail
	mkdir $path/mail
	rsync -av --quiet api/mail/ $path/mail --exclude mailer.go

	# .gitignore
	printf 'node_modules\n.env\n.env.local\n.env.*.local\n.idea\n.vscode\n*.suo\n*.ntvs*\n*.njsproj\n*.sln\n*.sw?\n.DS_Store\n' > $path/.gitignore

	#keep

	echo "Build for "$os" completed"
}

# Build for
build mac
build linux
build windows

# Commit
echo "Commiting & pushing build to GitHub..."
pwd
cd ./build
pwd
git add .
git commit -m "$commitmsg"
git push origin main