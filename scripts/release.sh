#!/bin/bash
# Extracts the release notes for a specific version from CHANGELOG.md

set -e

VERSION=$1

if [ -z "$VERSION" ]; then
    echo "Usage: $0 <version>"
    exit 1
fi

# Remove 'v' prefix if present to match CHANGELOG format
VERSION_CLEAN=${VERSION#v}

echo "Extracting release notes for version ${VERSION_CLEAN}..."

# Use awk to extract the block starting with the version header until the next version header
awk -v ver="[${VERSION_CLEAN}]" '
  $0 ~ "^## " ver {flag=1; print; next}
  /^## \[/ && flag {flag=0; exit}
  flag {print}
' CHANGELOG.md > RELEASE_NOTES.md

if [ ! -s RELEASE_NOTES.md ]; then
    echo "Warning: No release notes found for ${VERSION_CLEAN} in CHANGELOG.md"
    echo "## ${VERSION}" > RELEASE_NOTES.md
    echo "No release notes available." >> RELEASE_NOTES.md
fi

echo "Release notes written to RELEASE_NOTES.md"
