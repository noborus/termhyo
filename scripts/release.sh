#!/bin/bash

# Release script for termhyo
# Usage: ./scripts/release.sh v0.1.0

set -e

if [ $# -eq 0 ]; then
    echo "Usage: $0 <version>"
    echo "Example: $0 v0.1.0"
    exit 1
fi

VERSION=$1

# Validate version format
if [[ ! $VERSION =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo "Error: Version must be in format vX.Y.Z (e.g., v0.1.0)"
    exit 1
fi

echo "Preparing release $VERSION..."

# Check if we're on main branch
CURRENT_BRANCH=$(git branch --show-current)
if [ "$CURRENT_BRANCH" != "main" ]; then
    echo "Error: Must be on main branch for release. Current branch: $CURRENT_BRANCH"
    exit 1
fi

# Check if working directory is clean
if [ -n "$(git status --porcelain)" ]; then
    echo "Error: Working directory is not clean. Please commit or stash changes."
    git status --short
    exit 1
fi

# Pull latest changes
echo "Pulling latest changes..."
git pull origin main

# Run release checks
echo "Running release checks..."
make release-check

# Check if tag already exists
if git rev-parse "$VERSION" >/dev/null 2>&1; then
    echo "Error: Tag $VERSION already exists"
    exit 1
fi

# Update go.mod if needed
echo "Checking go.mod..."
go mod tidy

# Commit any go.mod changes
if [ -n "$(git status --porcelain go.mod go.sum)" ]; then
    echo "Updating go.mod and go.sum..."
    git add go.mod go.sum
    git commit -m "Update go.mod for release $VERSION"
fi

# Create and push tag
echo "Creating tag $VERSION..."
git tag -a "$VERSION" -m "Release $VERSION"

echo "Pushing tag to origin..."
git push origin "$VERSION"

echo "Release $VERSION completed successfully!"
echo "Visit https://github.com/noborus/termhyo/releases/tag/$VERSION to create the GitHub release"
