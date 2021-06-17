#!/bin/bash
#
# release.sh
#

# Set variables
version=$1
message=$2

# Check goreleaser passed
goreleaser check

# Check version is not empty
if [[ $version == "" ]]
  then
    echo "Add Version number"
    exit
fi

# Check commit message is not empty
if [[ $message == "" ]]
  then
    echo "Add commit message"
    exit
fi

rm -rf dist

git tag -a $version -m $message
git push origin $version

# Run goreleaser
goreleaser --rm-dist



