#!/bin/bash
#
# release.sh
#

# Set variables
version=$( cat VERSION )
message=$1

# Check if updater has been updated
read -p "Have you updated VERSION file?" -n 1 -r
echo    # (optional) move to a new line

if [[ $REPLY =~ ^[Yy]$ ]]
then
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

  git tag -a $version -m $message
  git push origin $version

  # Run goreleaser
  goreleaser release --rm-dist
fi





