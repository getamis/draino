#!/bin/bash
set -euo pipefail

VERSION=$(git describe --exact-match --tags 2> /dev/null || git rev-parse --short HEAD)
echo "Building version ${VERSION}"
docker build --tag "quay.io/amis/draino:${VERSION}" .
