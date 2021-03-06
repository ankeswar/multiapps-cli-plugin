#!/usr/bin/env bash

component=$1

version=$(cat VERSION)
major=$(echo $version | cut -d'.' -f 1)
minor=$(echo $version | cut -d'.' -f 2)
patch=$(echo $version | cut -d'.' -f 3)

case "$component" in
  major )
    major=$(expr $major + 1)
    minor=0
    patch=0
    ;;
  minor )
    minor=$(expr $minor + 1)
    patch=0
    ;;
  patch )
    patch=$(expr $patch + 1)
    ;;
  * )
    echo "Error - argument must be 'major', 'minor' or 'patch'"
    echo "Usage: bump-version [major | minor | patch]"
    exit 1
    ;;
esac

version=$major.$minor.$patch

echo "Updating VERSION file to $version..."
echo $version > VERSION

echo "Committing change..."
git reset .
git add VERSION
git commit -m "Bump version to $version"

echo "Creating v$version tag..."
git tag v$version
