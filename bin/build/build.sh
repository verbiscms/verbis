#!/bin/bash
#
# build.sh
#
# How to use:
# Run ./bin/build/build.sh from the root to get the right paths.
# Make sure to pass through a commit message to push the
# changes to the TryVerbis repo, for example:
#

# Set variables
path=./build

# Remove and create the build directory
cd ./build || exit
cd ..

function build() {

	# Build Vue
	cd admin || exit
	npm run build
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
	mkdir $path/storage && touch $path/storage/.gitkeep
	mkdir $path/storage/fields && touch $path/storage/fields/.gitkeep
	mkdir $path/storage/uploads
	mkdir $path/storage/dumps && touch $path/storage/dumps/.gitkeep
	mkdir $path/storage/forms && touch $path/storage/forms/.gitkeep

	# .gitignore
	printf 'node_modules\n.env\n.env.local\n.env.*.local\n.idea\n.vscode\n*.suo\n*.ntvs*\n*.njsproj\n*.sln\n*.sw?\n.DS_Store\n' > $path/.gitignore

}

# Build for
build