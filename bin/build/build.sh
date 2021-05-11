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
cd ./build || exit
git rm -rf .
git clean -fxd
cd ..

function build() {

	if [[ $commitmsg == "" ]]
		then
			echo "Add commit message"
			exit
	fi

	# Get the build path
	path=./build

	# Exec
  echo "Building executable..."
  CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o "./build/verbis-mac" -ldflags="-X 'github.com/ainsleyclark/verbis/api.SuperAdminString=false'"
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "./build/verbis-linux" -ldflags="-X 'github.com/ainsleyclark/verbis/api.SuperAdminString=false'"
  CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o "./build/verbis-windows.exe" -ldflags="-X 'github.com/ainsleyclark/verbis/api.SuperAdminString=false'"

	# Build Vue
	cd admin || exit
	npm run --silent --no-progress build
	cd ../

	# Build Web

	cd api/web || exit
	npm run --silent --no-progress prod
	cd ../../

	# Env
	cp .env.example $path/.env.example

	# API (Database)
	mkdir $path/api/

	# API (Web)
	mkdir $path/api/web && rsync -av --quiet api/web/ build/$os/api/web/ --exclude node_modules --exclude src --exclude package.json --exclude package-lock.json

	# Theme
	mkdir $path/themes && cp -a themes/Verbis $path/themes/

	# Admin
	mkdir $path/admin
	rsync -av --quiet admin/dist/ build/admin/

	# Storage
	mkdir $path/storage && touch $path/storage/.keep
	mkdir $path/storage/fields && touch $path/storage/fields/.keep
	mkdir $path/storage/uploads
	mkdir $path/storage/dumps && touch $path/storage/dumps/.keep
	mkdir $path/storage/logs && touch $path/storage/logs/.keep
	mkdir $path/storage/forms && touch $path/storage/forms/.keep

	# Mail
	#mkdir $path/mail
	#rsync -av --quiet api/mail/ $path/mail --exclude mailer.go

	# .gitignore
	printf 'node_modules\n.env\n.env.local\n.env.*.local\n.idea\n.vscode\n*.suo\n*.ntvs*\n*.njsproj\n*.sln\n*.sw?\n.DS_Store\n' > $path/.gitignore

	#keep

	echo "Build for "$os" completed"
}

# Build for
build

# Commit
echo "Commiting & pushing build to GitHub..."
pwd
cd ./build || exit
pwd
git add .
git commit -m "$commitmsg"
git push origin main