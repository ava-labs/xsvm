#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

if ! [[ "$0" =~ scripts/build.sh ]]; then
  echo "must be run from repository root"
  exit 255
fi

# Set default binary directory location
name="v3m4wPxaHpvGr8qfMeyK6PRW3idZrPHmYcMTt7oXdK47yurVH"

# Build xsvm, which can be run as a subprocess or as a CLI
mkdir -p ./build

echo "Building xsvm into ./build/"
go build -o ./build/$name ./cmd/xsvm/
go build -o ./build/xsvm ./cmd/xsvm/